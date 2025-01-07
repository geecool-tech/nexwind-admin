<template>
  <div class="container">
    <Breadcrumb :items="['权限管理', '管理员列表']" />
    <a-modal
      title="配置角色"
      :visible="showRole"
      @close="showRole = false"
      @cancel="showRole = false"
      @ok="handleRoleSave"
    >
      <div style="padding: 30px">
        <a-select v-model="curUserRole" multiple>
          <a-option
            v-for="(item, idx) in allRoleList"
            :key="idx"
            :disabled="item.id === 2"
            :value="item.id"
          >
            {{ item.name }}
          </a-option>
        </a-select>
      </div>
    </a-modal>
    <a-modal
      :width="500"
      :visible="showCreate"
      :title="formMode === 1 ? '添加管理员' : '编辑管理员'"
      @close="
        () => {
          showCreate = false;
        }
      "
      @cancel="
        () => {
          showCreate = false;
        }
      "
      @ok="handleFormSubmit"
    >
      <a-form ref="createForm" :model="formModel" label-align="right">
        <a-form-item
          :validate-trigger="['change', 'input']"
          field="name"
          :rules="[{ required: true, message: '昵称不能为空' }]"
          label="昵称"
        >
          <a-input v-model="formModel.name" placeholder="昵称"></a-input>
        </a-form-item>
        <a-form-item
          :validate-trigger="['change', 'input']"
          field="account"
          :rules="[{ required: true, message: '账号不能为空' }]"
          label="账号"
        >
          <a-input
            v-model="formModel.account"
            placeholder="请输入账号"
          ></a-input>
        </a-form-item>
        <a-form-item label="头像">
          <a-upload
            :action="proxy.$uploadOpt.url"
            :headers="proxy.$uploadOpt.headers"
            list-type="picture-card"
            :file-list="fileList"
            :show-upload-button="true"
            :show-file-list="false"
            @change="uploadChange"
          >
            <template #upload-button>
              <a-avatar :size="100" class="info-avatar">
                <template #trigger-icon>
                  <icon-camera />
                </template>
                <img v-if="fileList.length" :src="fileList[0].url" />
              </a-avatar>
            </template>
          </a-upload>
        </a-form-item>
        <a-form-item
          :validate-trigger="['change', 'input']"
          field="jobName"
          :rules="[{ required: true, message: '职位不能为空' }]"
          label="职位"
        >
          <a-input
            v-model="formModel.jobName"
            placeholder="请输入职位"
          ></a-input>
        </a-form-item>
        <a-form-item
          field="phone"
          :rules="[{ required: true, message: '手机号不能为空' }]"
          label="手机号"
        >
          <a-input
            v-model="formModel.phone"
            placeholder="请输入手机号"
          ></a-input>
        </a-form-item>
        <a-form-item
          field="organizationName"
          :rules="[{ required: true, message: '机构名称' }]"
          label="机构名称"
        >
          <a-input
            v-model="formModel.organizationName"
            placeholder="请输入机构名称"
          ></a-input>
        </a-form-item>
        <a-form-item
          :validate-trigger="['change', 'input']"
          field="email"
          label="邮箱"
          :rules="[{ required: true, message: '邮箱不能为空' }]"
        >
          <a-input v-model="formModel.email" placeholder="请输入邮箱"></a-input>
        </a-form-item>
        <a-form-item
          v-if="formMode === 1"
          :validate-trigger="['change', 'input']"
          field="password"
          label="密码"
          :rules="[{ required: true, message: '密码不能为空' }]"
        >
          <a-input
            v-model="formModel.password"
            :placeholder="formMode === 1 ? '请输入密码' : '留空则不修改'"
          ></a-input>
        </a-form-item>
        <a-form-item v-if="formMode === 2" field="password" label="密码">
          <a-input
            v-model="formModel.password"
            :placeholder="formMode === 1 ? '请输入密码' : '留空则不修改'"
          ></a-input>
        </a-form-item>
      </a-form>
    </a-modal>
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
                <a-form-item field="number" label="昵称">
                  <a-input v-model="conditions.name" placeholder="昵称" />
                </a-form-item>
              </a-col>
              <a-col :span="8">
                <a-form-item field="account" label="账号">
                  <a-input v-model="conditions.account" placeholder="账号" />
                </a-form-item>
              </a-col>
              <a-col :span="8">
                <a-form-item field="number" label="邮箱">
                  <a-input v-model="conditions.email" placeholder="邮箱" />
                </a-form-item>
              </a-col>
              <a-col :span="8">
                <a-form-item field="number" label="手机号">
                  <a-input v-model="conditions.phone" placeholder="手机号" />
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
            <a-button type="primary" @click="showCreateForm">
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
        <template #avatar="{ record }">
          <a-avatar>
            <img alt="avatar" :src="record.avatar" />
          </a-avatar>
        </template>
        <template #status="{ record }">
          <a-badge v-if="record.status === 1" text="启用中" status="success" />
        </template>
        <template #operations="{ record }">
          <a-button type="text" size="small" @click="showEdit(record)">
            <template #icon>
              <icon-edit />
            </template>
            编辑
          </a-button>

          <a-button type="text" size="small" @click="showRoleForm(record)">
            <template #icon>
              <icon-settings />
            </template>
            设置角色
          </a-button>
          <a-popconfirm
            v-if="record.accountId !== accountId"
            :content="'确定删除账号 [' + record.account + ']?'"
            @ok="del(record)"
          >
            <a-button type="text" size="small">
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
  import {
    computed,
    ref,
    reactive,
    watch,
    nextTick,
    getCurrentInstance,
  } from 'vue';
  import { useI18n } from 'vue-i18n';
  import { Message } from '@arco-design/web-vue';
  import useLoading from '@/hooks/loading';
  import {
    queryAdminList,
    AdminListRecord,
    AdminListParam,
    delAdmin,
    saveAdmin,
    queryUserRole,
    setUserRole,
    RoleListRecord,
    queryRoleList,
  } from '@/api/auth';
  import { Pagination } from '@/types/global';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import Sortable from 'sortablejs';
  import type {
    FileItem,
    RequestOption,
  } from '@arco-design/web-vue/es/upload/interfaces';
  import { userUploadApi } from '@/api/user-center';
  import { useUserStore } from '@/store';

  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  type Column = TableColumnData & { checked?: true };

  const generateFormModel = () => {
    return {
      id: 0,
      accountId: '',
      name: '',
      avatar: '',
      email: '',
      account: '',
      organizationName: '',
      jobName: '',
      phone: '',
      status: 1,
      registrationDate: '',
      password: '',
    };
  };
  const userStore = useUserStore();
  const file = {
    uid: '-2',
    name: userStore.avatar,
    url: userStore.avatar,
  };
  const { accountId } = userStore;
  const conditions = ref<AdminListParam>({
    current: 1,
    pageSize: 10,
    name: '',
    email: '',
    nickname: '',
    account: '',
    mobile: '',
  });
  const { loading, setLoading } = useLoading(true);
  const { t } = useI18n();
  const { proxy } = getCurrentInstance();
  const renderData = ref<AdminListRecord[]>([]);
  const formModel = ref(generateFormModel());
  const cloneColumns = ref<Column[]>([]);
  const fileList = ref<FileItem[]>([file]);
  const showColumns = ref<Column[]>([]);
  const formMode = ref(1);
  const size = ref<SizeProps>('medium');
  const showCreate = ref<boolean>(false);
  const basePagination: Pagination = {
    current: 1,
    pageSize: 20,
  };

  const pagination = reactive({
    ...basePagination,
  });
  const createForm = ref();
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
  const customRequest = (options: RequestOption) => {
    // docs: https://axios-http.com/docs/cancellation
    const controller = new AbortController();

    (async function requestWrap() {
      const {
        onProgress,
        onError,
        onSuccess,
        fileItem,
        name = 'file',
      } = options;
      onProgress(20);
      const formData = new FormData();
      formData.append(name as string, fileItem.file as Blob);
      const onUploadProgress = (event: ProgressEvent) => {
        let percent;
        if (event.total > 0) {
          percent = (event.loaded / event.total) * 100;
        }
        onProgress(parseInt(String(percent), 10), event);
      };

      try {
        // https://github.com/axios/axios/issues/1630
        // https://github.com/nuysoft/Mock/issues/127

        const res = await userUploadApi(formData, {
          controller,
          onUploadProgress,
        });
        onSuccess(res);
      } catch (error) {
        onError(error);
      }
    })();
    return {
      abort() {
        controller.abort();
      },
    };
  };
  const columns = computed<TableColumnData[]>(() => [
    {
      title: t('searchTable.columns.index'),
      dataIndex: 'id',
    },
    {
      title: 'AccountId',
      dataIndex: 'accountId',
      width: 120,
    },
    {
      title: '昵称',
      dataIndex: 'name',
      width: 140,
    },
    {
      title: '头像',
      dataIndex: 'avatar',
      slotName: 'avatar',
    },
    {
      title: '账号',
      dataIndex: 'account',
    },
    {
      title: '机构名称',
      dataIndex: 'organizationName',
      width: 100,
    },
    {
      title: '职位',
      dataIndex: 'jobName',
    },
    {
      title: '手机号',
      dataIndex: 'phone',
    },
    {
      title: '邮箱',
      dataIndex: 'email',
    },
    {
      title: '状态',
      dataIndex: 'status',
      slotName: 'status',
      width: 90,
    },
    {
      title: '注册时间',
      dataIndex: 'registrationDate',
    },
    {
      title: t('searchTable.columns.operations'),
      dataIndex: 'operations',
      fixed: 'right',
      width: 200,
      slotName: 'operations',
    },
  ]);
  const curUserRole = ref<number[]>([]);
  const allRoleList = ref<RoleListRecord>([]);

  const fetchRoleList = async () => {
    const resp = await queryRoleList({ current: 1, pageSize: 100 });
    allRoleList.value = resp.data.list;
  };
  fetchRoleList();
  const showRole = ref(false);
  const prevUserRole = ref([]);
  const handleRoleSave = async () => {
    const added = curUserRole.value.filter(
      (item) => !prevUserRole.value.includes(item)
    );
    const removed = prevUserRole.value.filter(
      (item) => !curUserRole.value.includes(item)
    );
    const resp = await setUserRole({
      uid: formModel.value.id,
      add: added,
      remove: removed,
    });
    if (resp.code === 20000) {
      Message.success('配置成功~');
      showRole.value = false;
    } else {
      Message.error('配置失败~');
      showRole.value = false;
    }
  };
  const showRoleForm = async (record: any) => {
    const resp = await queryUserRole(record);
    formModel.value = record;
    if (resp.code === 20000) {
      curUserRole.value = resp.data;
      prevUserRole.value = resp.data;
      showRole.value = true;
    }
  };
  const del = (record: any) => {
    delAdmin(record.id).then((res: any) => {
      if (res?.code === 20000) {
        Message.success('删除成功');
        fetchData();
      } else {
        Message.error(res.message);
      }
    });
  };
  const originFile=ref({
    uid: -1,
    name: 'avatar.png',
    url: '',
  });
  const showEdit = (record: any) => {
    formMode.value = 2;
    fileList.value = [
      {
        uid: record.id,
        name: 'avatar.png',
        url: record.avatar,
      },
    ];
    showCreate.value = true;
    formModel.value = cloneDeep(record);
  };
  const handleFormSubmit = () => {
    createForm.value.validate(async (valid: any) => {
      if (valid) {
        return;
      }
      await saveAdmin(formModel.value).then((res: any) => {
        if (res?.code === 20000) {
          showCreate.value = false;
          formModel.value = generateFormModel();
          Message.success(formMode.value === 1 ? '添加成功' : '保存成功');
          fetchData();
        } else {
          Message.error(res.message);
        }
      });
    });
  };
  const uploadChange = (fileItemList: FileItem[], fileItem: FileItem) => {
    if (fileItem.status === 'done') {
      formModel.value.avatar = fileItem.response.data.url;
    }
    console.log(formModel.value);
    fileList.value = [fileItem];
  };
  const showCreateForm = () => {
    formMode.value = 1;
    fileList.value=[originFile.value];
    showCreate.value = true;
  };
  const fetchData = async (
    params: AdminListParam = {
      current: 1,
      pageSize: 20,
      name: '',
      nickname: '',
      account: '',
      mobile: '',
      email: '',
    }
  ) => {
    setLoading(true);
    try {
      const { data } = await queryAdminList(params);
      renderData.value = data.list;
      pagination.current = params.current;
      pagination.total = data.total;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const search = () => {
    fetchData({
      ...basePagination,
      ...conditions.value,
    } as unknown as AdminListParam);
  };
  const onPageChange = (current: number) => {
    fetchData({
      account: '',
      email: '',
      mobile: '',
      name: '',
      nickname: '',
      ...basePagination,
      current,
    });
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
