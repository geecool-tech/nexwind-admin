<template>
  <div class="container">
    <a-modal
      ref="createForm"
      :visible="showCreate"
      :title="createMode === 1 ? '添加' : '编辑'"
      @ok="handleSubmit"
      @cancel="showCreate = false"
      @close="showCreate = false"
    >
      <a-form ref="createForm" :model="formModel">
        <a-form-item
          label="标题"
          field="title"
          :rules="[{ required: true, message: '标题不能为空' }]"
        >
          <a-input v-model="formModel.title"></a-input>
        </a-form-item>
        <a-form-item
          label="组件名称"
          field="name"
          :rules="[{ required: true, message: '组件名称' }]"
        >
          <a-input v-model="formModel.name"></a-input>
        </a-form-item>
        <a-form-item label="icon" field="icon">
          <a-input v-model="formModel.meta.icon"></a-input>
        </a-form-item>
        <a-form-item
          label="path"
          field="path"
          :rules="[{ required: true, message: 'path' }]"
        >
          <a-input v-model="formModel.path"></a-input>
        </a-form-item>
      </a-form>
    </a-modal>
    <Breadcrumb :items="['权限管理', '菜单列表']" />
    <a-card class="general-card" title="菜单列表">
      <a-row style="margin-bottom: 16px">
        <a-col :span="12">
          <a-space>
            <a-button type="primary" @click="addRoot">
              <template #icon>
                <icon-plus />
              </template>
              添加根目录
            </a-button>
          </a-space>
        </a-col>
        <a-col
          :span="12"
          style="display: flex; align-items: center; justify-content: end"
        >
          <a-tooltip :content="$t('searchTable.actions.refresh')">
            <div class="action-icon" @click="search"
              ><icon-refresh size="18"
            /></div>
          </a-tooltip>
          <a-dropdown @select="handleSelectDensity">
            <a-tooltip :content="$t('searchTable.actions.density')">
              <div class="action-icon"><icon-line-height size="18" /></div>
            </a-tooltip>
            <template #content>
              <a-doption
                v-for="item in densityList"
                :key="item.value"
                :value="item.value"
                :class="{ active: item.value === size }"
              >
                <span>{{ item.name }}</span>
              </a-doption>
            </template>
          </a-dropdown>
          <a-tooltip :content="$t('searchTable.actions.columnSetting')">
            <a-popover
              trigger="click"
              position="bl"
              @popup-visible-change="popupVisibleChange"
            >
              <div class="action-icon"><icon-settings size="18" /></div>
              <template #content>
                <div id="tableSetting">
                  <div
                    v-for="(item, index) in showColumns"
                    :key="item.dataIndex"
                    class="setting"
                  >
                    <div style="margin-right: 4px; cursor: move">
                      <icon-drag-arrow />
                    </div>
                    <div>
                      <a-checkbox
                        v-model="item.checked"
                        @change="
                          handleChange($event, item as TableColumnData, index)
                        "
                      >
                      </a-checkbox>
                    </div>
                    <div class="title">
                      {{ item.title === '#' ? '序列号' : item.title }}
                    </div>
                  </div>
                </div>
              </template>
            </a-popover>
          </a-tooltip>
        </a-col>
      </a-row>
      <a-table
        row-key="id"
        :loading="loading"
        :pagination="false"
        :columns="(cloneColumns as TableColumnData[])"
        :data="renderData"
        :bordered="false"
        :size="size"
      >
        <template #index="{ rowIndex }">
          {{ rowIndex + 1 + (pagination.current - 1) * pagination.pageSize }}
        </template>
        <template #title="{ record }">
          <span v-if="record.meta.icon">
            <component :is="record.meta.icon" />
          </span>
          {{ record.title }}
        </template>
        <template #type="{ record }">
          <a-tag v-if="record.type === 1" color="cyan">目录</a-tag>
          <a-tag v-if="record.type === 2" color="arcoblue">菜单</a-tag>
          <a-tag v-if="record.type === 3" color="red">操作</a-tag>
        </template>
        <template #contentType="{ record }">
          <a-space>
            <a-avatar
              v-if="record.contentType === 'img'"
              :size="16"
              shape="square"
            >
              <img
                alt="avatar"
                src="//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/581b17753093199839f2e327e726b157.svg~tplv-49unhts6dw-image.image"
              />
            </a-avatar>
            <a-avatar
              v-else-if="record.contentType === 'horizontalVideo'"
              :size="16"
              shape="square"
            >
              <img
                alt="avatar"
                src="//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/77721e365eb2ab786c889682cbc721c1.svg~tplv-49unhts6dw-image.image"
              />
            </a-avatar>
            <a-avatar v-else :size="16" shape="square">
              <img
                alt="avatar"
                src="//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/ea8b09190046da0ea7e070d83c5d1731.svg~tplv-49unhts6dw-image.image"
              />
            </a-avatar>
            {{ $t(`searchTable.form.contentType.${record.contentType}`) }}
          </a-space>
        </template>

        <template #operations="{ record }">
          <a-button type="text" size="small" @click="edit(record)">
            <template #icon>
              <icon-edit />
            </template>
            编辑
          </a-button>
          <a-button
            v-if="record.type === 1"
            type="text"
            size="small"
            @click="addChild(record)"
          >
            <template #icon>
              <icon-drive-file />
            </template>
            添加子菜单
          </a-button>
          <a-button
            v-if="record.type === 2"
            type="text"
            size="small"
            @click="addOperation(record)"
          >
            <template #icon>
              <icon-copy />
            </template>
            添加操作
          </a-button>
          <a-popconfirm
            :content="'确定删除 [ ' + record.title + ' ]?'"
            @ok="del(record.id)"
          >
            <a-button type="text" size="small">
              <template #icon>
                <icon-delete />
              </template>
              删除
            </a-button>
          </a-popconfirm>
          <a-button
            v-if="record.type === 1"
            type="text"
            size="small"
            @click="addMenu(record)"
          >
            <template #icon>
              <icon-folder />
            </template>
            添加目录
          </a-button>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
  import { useAppStore } from '@/store';
  import { Message } from '@arco-design/web-vue';
  import { computed, nextTick, reactive, ref, watch } from 'vue';
  import { useI18n } from 'vue-i18n';
  import useLoading from '@/hooks/loading';
  import {
    MenuListRecord,
    saveMenu,
    MenuMeta,
    delMenu,
    queryMenuList,
  } from '@/api/auth';
  import { Pagination } from '@/types/global';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import Sortable from 'sortablejs';

  const appStore = useAppStore();
  const generateFormModel = (
    type?: number,
    parentId?: number | null
  ): MenuListRecord => {
    let meta: MenuMeta = {
      requiresAuth: false,
    };
    if (type === 1) {
      meta = {
        requiresAuth: true,
        icon: '',
      };
    }
    if (type === 2) {
      meta = {
        requiresAuth: true,
        hideChildrenInMenu: true,
        icon: '',
      };
    }
    if (type === 3) {
      meta = {
        hideInMenu: true,
        requiresAuth: true,
        icon: '',
      };
    }
    return {
      title: '',
      name: '',
      parent_id: parentId,
      path: '',
      type: type || 1,
      meta,
    };
  };
  const createForm = ref();
  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  type Column = TableColumnData & { checked?: true };
  const showCreate = ref(false);
  const { loading, setLoading } = useLoading(true);
  const { t } = useI18n();
  const createMode = ref(1);
  const renderData = ref<MenuListRecord[]>([]);
  const formModel = ref(generateFormModel());
  const cloneColumns = ref<Column[]>([]);
  const showColumns = ref<Column[]>([]);
  const size = ref<SizeProps>('medium');

  const basePagination: Pagination = {
    current: 1,
    pageSize: 20,
  };

  const pagination = reactive({
    ...basePagination,
  });
  const densityList = computed(() => [
    {
      name: t('searchTable.size.mini'),
      value: 'mini',
    },
    {
      name: t('searchTable.size.small'),
      value: 'small',
    },
    {
      name: t('searchTable.size.medium'),
      value: 'medium',
    },
    {
      name: t('searchTable.size.large'),
      value: 'large',
    },
  ]);
  const edit = (record: MenuListRecord) => {
    createMode.value = 2;
    formModel.value = cloneDeep(record);
    showCreate.value = true;
  };
  const columns = computed<TableColumnData[]>(() => [
    {
      title: t('searchTable.columns.index'),
      dataIndex: 'id',
    },
    {
      title: '菜单标题',
      dataIndex: 'title',
      slotName: 'title',
    },
    {
      title: '类型',
      dataIndex: 'type',
      slotName: 'type',
    },
    {
      title: 'path',
      dataIndex: 'path',
    },
    {
      title: '组件',
      dataIndex: 'name',
    },
    {
      title: '创建时间',
      width: 170,
      dataIndex: 'created_time',
    },
    {
      title: t('searchTable.columns.operations'),
      dataIndex: 'operations',
      fixed: 'right',
      width: 240,
      slotName: 'operations',
    },
  ]);
  const addChild = (record: MenuListRecord) => {
    formModel.value = generateFormModel(2, record.id);
    createMode.value = 1;
    showCreate.value = true;
  };
  const addOperation = (record: MenuListRecord) => {
    formModel.value = generateFormModel(3, record.id);
    createMode.value = 1;
    showCreate.value = true;
  };
  const addMenu = (record: MenuListRecord) => {
    formModel.value = generateFormModel(1, record.id);
    createMode.value = 1;
    showCreate.value = true;
  };
  const handleSubmit = () => {
    createForm.value.validate(async (valid: any) => {
      if (valid) {
        return;
      }
      await saveMenu(formModel.value).then((res: any) => {
        if (res?.code === 20000) {
          showCreate.value = false;
          formModel.value = generateFormModel();
          Message.success(createMode.value === 1 ? '添加成功' : '保存成功');
          appStore.fetchServerMenuConfig();
          fetchData();
        } else {
          Message.error(res.message);
        }
      });
    });
  };
  const addRoot = () => {
    showCreate.value = true;
    formModel.value = generateFormModel(1, null);
  };
  const del = async (id: number) => {
    delMenu(id).then((res: any) => {
      if (res.code === 20000) {
        Message.success('删除成功');
        appStore.fetchServerMenuConfig();
        fetchData();
      } else {
        Message.error(res.message);
      }
    });
  };
  const fetchData = async () => {
    setLoading(true);
    try {
      const { data } = await queryMenuList();
      renderData.value = data.list;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const search = () => {
    fetchData();
  };

  fetchData();
  const reset = () => {
    formModel.value = generateFormModel();
  };

  const handleSelectDensity = (
    val: string | number | Record<string, any> | undefined,
    e: Event
  ) => {
    size.value = val as SizeProps;
  };

  const handleChange = (
    checked: boolean | (string | boolean | number)[],
    column: Column,
    index: number
  ) => {
    if (!checked) {
      cloneColumns.value = showColumns.value.filter(
        (item) => item.dataIndex !== column.dataIndex
      );
    } else {
      cloneColumns.value.splice(index, 0, column);
    }
  };

  const exchangeArray = <T extends Array<any>>(
    array: T,
    beforeIdx: number,
    newIdx: number,
    isDeep = false
  ): T => {
    const newArray = isDeep ? cloneDeep(array) : array;
    if (beforeIdx > -1 && newIdx > -1) {
      // 先替换后面的，然后拿到替换的结果替换前面的
      newArray.splice(
        beforeIdx,
        1,
        newArray.splice(newIdx, 1, newArray[beforeIdx]).pop()
      );
    }
    return newArray;
  };
  const popupVisibleChange = (val: boolean) => {
    if (val) {
      nextTick(() => {
        const el = document.getElementById('tableSetting') as HTMLElement;
        const sortable = new Sortable(el, {
          onEnd(e: any) {
            const { oldIndex, newIndex } = e;
            exchangeArray(cloneColumns.value, oldIndex, newIndex);
            exchangeArray(showColumns.value, oldIndex, newIndex);
          },
        });
      });
    }
  };

  watch(
    () => columns.value,
    (val) => {
      cloneColumns.value = cloneDeep(val);
      cloneColumns.value.forEach((item, index) => {
        item.checked = true;
      });
      showColumns.value = cloneDeep(cloneColumns.value);
    },
    { deep: true, immediate: true }
  );
</script>

<script lang="ts">
  export default {
    name: 'MenuList',
  };
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
  }
  :deep(.arco-table-th) {
    &:last-child {
      .arco-table-th-item-title {
        margin-left: 16px;
      }
    }
  }
  .action-icon {
    margin-left: 12px;
    cursor: pointer;
  }
  .active {
    color: #0960bd;
    background-color: #e3f4fc;
  }
  .setting {
    display: flex;
    align-items: center;
    width: 200px;
    .title {
      margin-left: 12px;
      cursor: pointer;
    }
  }
</style>
