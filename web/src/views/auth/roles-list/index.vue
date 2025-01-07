<template>
  <div class="container">
    <a-modal @ok="handleSubmit" @close="showCreate = false" @cancel="showCreate = false"  :title="operationMode===1?'新增角色':'编辑角色'" :visible="showCreate">
      <a-form ref="createForm" :model="formModel">
        <a-form-item :rules="[{ required: true, message: '名称不能为空' }]" label="名称" field="name">
          <a-input v-model="formModel.name"></a-input>
        </a-form-item>
        <a-form-item :rules="[{ required: true, message: '描述不能为空' }]" label="描述" field="description">
          <a-input v-model="formModel.description"></a-input>
        </a-form-item>
      </a-form>

    </a-modal>
    <a-drawer
      :width="600"
      @ok="showAuthNode=false"
      :visible="showAuthNode"
      :title="'设置[ ' + formModel.name + ' ]权限节点'"
      @cancel="showAuthNode = false"
      @close="showAuthNode = false"
    >

      <template #title></template>
      <div>
        <a-tabs default-active-key="1" type="card" size="large">
          <a-tab-pane key="1" type="card" >
            <template #title>菜单权限</template>
            <div style="padding: 15px">
              <a-tree
                  @check="handleMenuCheck"
                  @change="handleChangeMenu"
                  :check-strictly="true"
                v-model:checked-keys="menuCheckedKeys"
                :field-names="{
                  key: 'id',
                  title: 'title',
                  children: 'children',
                }"
                :checkable="true"
                :data="authNode.menu_list"
              />
            </div>
          </a-tab-pane>
          <a-tab-pane key="2" >
            <template #title>Api 权限</template>
            <div style="padding: 15px">
              <a-tree
                  blockNode
                  @check="handleApiCheck"
                  v-model:checked-keys="apiCheckedKeys"
                  :field-names="{
                  key: 'id',
                  title: 'title',
                  path:'path',
                  children: 'children',
                }"
                  :checkable="true"
                  :data="authNode.tree"
              >
                <template #title="{ title,path }">
                  <div style="display: flex;gap: 50px">
                    <div>{{title?title:'未命名接口'}}</div>
                    <div>{{path}}</div>
                  </div>
                </template>

              </a-tree>

            </div>
          </a-tab-pane>
        </a-tabs>
      </div>
    </a-drawer>
    <Breadcrumb :items="['权限管理', '角色列表']" />
    <a-card class="general-card" :title="$t('menu.list.searchTable')">
      <a-row>
        <a-col :flex="1">
          <a-form
            :model="formModel"
            :label-col-props="{ span: 6 }"
            :wrapper-col-props="{ span: 18 }"
            label-align="left"
          >
            <a-row :gutter="16">
              <a-col :span="8">
                <a-form-item field="number" label="角色名称">
                  <a-input v-model="formModel.name" placeholder="角色名称" />
                </a-form-item>
              </a-col>
            </a-row>
          </a-form>
        </a-col>
        <a-divider style="height: 84px" direction="vertical" />
        <a-col :flex="'86px'" style="text-align: right">
          <a-space direction="vertical" :size="18">
            <a-button type="primary" @click="search">
              <template #icon>
                <icon-search />
              </template>
              {{ $t('searchTable.form.search') }}
            </a-button>
            <a-button @click="reset">
              <template #icon>
                <icon-refresh />
              </template>
              {{ $t('searchTable.form.reset') }}
            </a-button>
          </a-space>
        </a-col>
      </a-row>
      <a-divider style="margin-top: 0" />
      <a-row style="margin-bottom: 16px">
        <a-col :span="12">
          <a-space>
            <a-button @click="setShowCreate" type="primary">
              <template #icon>
                <icon-plus />
              </template>
              {{ $t('searchTable.operation.create') }}
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
        :pagination="pagination"
        :columns="(cloneColumns as TableColumnData[])"
        :data="renderData"
        :bordered="false"
        :size="size"
        @page-change="onPageChange"
      >
        <template #index="{ rowIndex }">
          {{ rowIndex + 1 + (pagination.current - 1) * pagination.pageSize }}
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
        <template #filterType="{ record }">
          {{ $t(`searchTable.form.filterType.${record.filterType}`) }}
        </template>
        <template #status="{ record }">
          <span v-if="record.status === 'offline'" class="circle"></span>
          <span v-else class="circle pass"></span>
          {{ $t(`searchTable.form.status.${record.status}`) }}
        </template>
        <template #operations="{ record }">
          <a-button @click="editRecord(record)" type="text" size="small">
            <template #icon>
              <icon-edit />
            </template>
            编辑
          </a-button>

          <a-button type="text" size="small" @click="showRoleAuth(record)">
            <template #icon>
              <icon-settings />
            </template>
            设置权限
          </a-button>
          <a-popconfirm @ok="del(record)" v-if="record.id!==1 && record.id!==2" :content="'确定删除角色 ['+record.name +']?'">
            <a-button  type="text" size="small" >
              <template #icon>
                <icon-delete />
              </template>
              删除
            </a-button>
          </a-popconfirm>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, reactive, watch, nextTick, onMounted } from 'vue';
  import { useI18n } from 'vue-i18n';
  import useLoading from '@/hooks/loading';
  import {
    queryRoleList,
    getAuthNode,
    RoleListRecord,
    RoleListParams,
    getRoleAccess,
    RoleAccessRes,
    authRoleMenu,
    authRoleApi,
    delRole,
    saveRole
  } from '@/api/auth';
  import { Pagination } from '@/types/global';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import Sortable from 'sortablejs';
  import {Message} from "@arco-design/web-vue";

  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  type Column = TableColumnData & { checked?: true };
  const authNode = ref<RoleAccessRes>({
    tree: [],
    api_list: {},
    menu_list: [],
  });
  const showAuthNode = ref(false);

  const generateFormModel = () => {
    const model: RoleListRecord = {
      id: 0,
      name: '',
      description: '',
      created_time: '',
      updated_time: '',
    };
    return model;
  };

  const menuCheckedKeys = ref([]);
  const prevMenuCheckedKeys = ref([]);
  const prevApiCheckedKeys = ref<number[]>([]);
  const apiCheckedKeys = ref([]);
  const { loading, setLoading } = useLoading(true);
  const { t } = useI18n();
  const renderData = ref<RoleListRecord[]>([]);
  const formModel = ref<RoleListRecord>(generateFormModel());
  const cloneColumns = ref<Column[]>([]);
  const showColumns = ref<Column[]>([]);
  const conditions = ref<RoleListParams>({
    name: '',
    pageSize: 20,
    current: 1,
  });
  const size = ref<SizeProps>('medium');

  const basePagination: Pagination = {
    current: 1,
    pageSize: 20,
  };
  const pagination = reactive({
    ...basePagination,
  });
  const showCreate = ref<boolean>(false);
  const createForm = ref();
  const operationMode = ref(1);
  const del=async (record:any)=>{
    const resp=await delRole(record.id);
    if (resp.code===20000){
      Message.success('删除成功');
      fetchData(conditions.value);
    }else {
      Message.error(resp.message);
    }
  }
  const setShowCreate = () => {
    formModel.value=generateFormModel();
    showCreate.value = true;
    operationMode.value=1;
  }
  const editRecord=(record:any)=>{
    formModel.value=record;
    operationMode.value=2;
    showCreate.value=true;
  };
  const handleSubmit=async ()=>{
    const resp=await saveRole(formModel.value);
    if (resp.code===20000){
      Message.success(operationMode.value===1?'创建成功':'修改成功');
      showCreate.value=false;
      fetchData();
    }else {
      Message.error(resp.message);
    }
  };
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
  const columns = computed<TableColumnData[]>(() => [
    {
      title: t('searchTable.columns.index'),
      dataIndex: 'id',
    },
    {
      title: '角色名称',
      dataIndex: 'name',
    },
    {
      title: '描述',
      dataIndex: 'description',
    },
    {
      title: '创建时间',
      dataIndex: 'created_time',
    },
    {
      title: '修改时间',
      dataIndex: 'updated_time',
    },
    {
      title: t('searchTable.columns.operations'),
      dataIndex: 'operations',
      fixed: 'right',
      width: 200,
      slotName: 'operations',
    },
  ]);

  const fetchAuthNode=async ()=>{
    const resp = await getAuthNode();
    authNode.value = resp.data;
  };
  const fetchData = async (conditions?: RoleListParams) => {
    setLoading(true);
    try {
      const { data } = await queryRoleList(conditions);
      renderData.value = data.list;
      pagination.current = conditions.current;
      pagination.total = data.total;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };
  const handleChangeMenu=()=>{};
  const handleSubmitAuth=()=>{};
  const showRoleAuth = async (record:any)=>{
    formModel.value = cloneDeep(record);

    const {data}=await getRoleAccess(record.id);
    menuCheckedKeys.value = data.menu;
    apiCheckedKeys.value = data.api;
    prevApiCheckedKeys.value=cloneDeep(apiCheckedKeys.value);
    prevMenuCheckedKeys.value=cloneDeep(menuCheckedKeys.value);
    showAuthNode.value = true;
  };
  const handleApiCheck=(checkedIds:any)=>{
    const added = apiCheckedKeys.value.filter(id => !prevApiCheckedKeys.value.includes(id));
    const removed = prevApiCheckedKeys.value.filter(id => !apiCheckedKeys.value.includes(id));
    authRoleApi({
      role_id:formModel.value.id,
      add:added.filter(item => typeof item === 'number'),
      remove:removed.filter(item => typeof item === 'number')
    }).then(res=>{
      if(res.code!==20000){
        Message.error("修改失败,请刷新页面");
      }
    })
    prevApiCheckedKeys.value = [...apiCheckedKeys.value];

  };
  const handleMenuCheck=(checkedIds:any)=>{
    const added = menuCheckedKeys.value.filter(id => !prevMenuCheckedKeys.value.includes(id));
    const removed = prevMenuCheckedKeys.value.filter(id => !menuCheckedKeys.value.includes(id));
    authRoleMenu({
      role_id:formModel.value.id,
      add:added,
      remove:removed
    }).then(res=>{
      if(res.code!==20000){
        Message.error("修改失败,请刷新页面");
      }
    })
    prevMenuCheckedKeys.value = [...checkedIds];

  };
  const search = () => {
    fetchData({
      ...basePagination,
      ...formModel.value,
    } as unknown as RoleListParams);
  };
  const onPageChange = (current: number) => {
    conditions.value.current = current;
    fetchData(conditions.value);
  };
  fetchAuthNode();
  fetchData(conditions.value);
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
    name: 'RolesList',
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
