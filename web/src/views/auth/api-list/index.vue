<template>
  <div class="container">
    <a-modal
      title="编辑 Api"
      :visible="showForm"
      @ok="handleSave"
      @close="showForm = false"
      @cancel="showForm = false"
    >
      <a-form ref="EditForm" :model="formModel">
        <a-form-item
          field="title"
          :validate-trigger="['change', 'input']"
          :rules="[{ required: true, message: '接口名称不能为空' }]"
          label="Api 名称"
        >
          <a-input v-model="formModel.title"></a-input>
        </a-form-item>
        <a-form-item label="path">
          <a-input v-model="formModel.path" disabled></a-input>
        </a-form-item>
        <a-form-item
          field="scope"
          label="Scope"
          :validate-trigger="['change', 'input']"
          :rules="[{ required: true, message: 'Scope 不能为空' }]"
        >
          <a-select
            v-model="formModel.scope"
            placeholder="请选择 Scope"
            allow-create
          >
            <a-option
              v-for="(scope, key) in apiListRecord?.all_scope"
              :key="key"
              :value="key"
              >{{ key }}</a-option
            >
          </a-select>
        </a-form-item>
        <a-form-item
          field="group"
          :validate-trigger="['change', 'input']"
          :rules="[{ required: true, message: '分组不能为空' }]"
          label="分组"
        >
          <a-select
            v-model="formModel.group"
            placeholder="请选择 Scope"
            allow-create
          >
            <a-option
              v-for="(group, key) in apiListRecord['all_scope']?.[
                formModel.scope
              ]"
              :key="key"
              :value="group"
              >{{ group }}</a-option
            >
          </a-select>
        </a-form-item>
        <a-form-item
          :validate-trigger="['change', 'input']"
          :rules="[{ required: true, message: '请求方式不能为空' }]"
          label="请求方式"
        >
          <a-select
            v-model="formModel.method"
            placeholder="请选择请求方式"
            allow-create
          >
            <a-option value="GET">GET</a-option>
            <a-option value="POST">POST</a-option>
            <a-option value="DELETE">DELETE</a-option>
            <a-option value="PUT">PUT</a-option>
          </a-select>
        </a-form-item>
        <a-form-item
          label="登录"
          :validate-trigger="['change', 'input']"
          :rules="[{ required: true, message: '是否登录不能为空' }]"
        >
          <a-select v-model="formModel.need_login" placeholder="请选择是否登录">
            <a-option :value="1">是</a-option>
            <a-option :value="2">否</a-option>
          </a-select>
        </a-form-item>
        <a-form-item
          label="鉴权"
          :validate-trigger="['change', 'input']"
          :rules="[{ required: true, message: '是否鉴权不能为空' }]"
        >
          <a-select v-model="formModel.need_auth" placeholder="请选择是否鉴权">
            <a-option :value="1">是</a-option>
            <a-option :value="2">否</a-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>
    <Breadcrumb :items="['权限管理', 'Api 列表']" />
    <a-card class="general-card" title="Api 列表">
      <div v-for="(scope, sKey) in apiListRecord?.formatted_list" :key="sKey">
        <div style="padding: 5px 0; font-size: 15px; color: black"
          ><icon-layers />
          {{ sKey }}
          <icon-caret-right
            v-if="!expandKeys[sKey]"
            style="cursor: pointer"
            @click="expandScope(sKey)"
          />
          <icon-caret-down
            v-if="expandKeys[sKey]"
            style="cursor: pointer"
            @click="expandScope(sKey)"
          />
        </div>
        <div v-if="expandKeys[sKey]" style="padding-left: 10px">
          <div v-for="(list, gKey) in scope" :key="gKey">
            <div style="padding: 5px; font-size: 15px; color: black">
              <icon-folder /> {{ gKey }}
              <icon-caret-right
                v-if="!expandKeys[sKey][gKey]"
                style="cursor: pointer"
                @click="expandGroup(sKey, gKey)"
              />
              <icon-caret-down
                v-if="expandKeys[sKey][gKey]"
                style="cursor: pointer"
                @click="expandGroup(sKey, gKey)"
              />
            </div>
            <div v-if="expandKeys[sKey][gKey]" style="padding: 15px">
              <a-table
                size="small"
                :columns="(cloneColumns as TableColumnData[])"
                row-key="id"
                :data="list"
              >
                <template #method="{ record }">
                  <a-tag v-if="record.method === 'GET'" color="green">{{
                    record.method
                  }}</a-tag>
                  <a-tag v-if="record.method === 'POST'" color="cyan">{{
                    record.method
                  }}</a-tag>
                  <a-tag v-if="record.method === 'HEAD'" color="arcoblue">{{
                    record.method
                  }}</a-tag>
                </template>
                <template #need_login="{ record }">
                  {{ record['need_login'] === 1 ? '是' : '否' }}
                </template>
                <template #need_auth="{ record }">
                  {{ record['need_auth'] === 1 ? '是' : '否' }}
                </template>
                <template #operation="{ record }">
                  <div style="display: flex">
                    <a-button
                      type="text"
                      size="small"
                      @click="showEdit(record)"
                    >
                      <template #icon>
                        <icon-edit />
                      </template>
                      编辑
                    </a-button>

<!--                    <a-popconfirm-->
<!--                      :content="'确定删除接口 [' + record.title + ']?'"-->
<!--                      @ok="del(record)"-->
<!--                    >-->
<!--                      <a-button type="text" size="small">-->
<!--                        <template #icon>-->
<!--                          <icon-delete />-->
<!--                        </template>-->
<!--                        删除-->
<!--                      </a-button>-->
<!--                    </a-popconfirm>-->
                  </div>
                </template>
              </a-table>
            </div>
          </div>
        </div>
      </div>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref, reactive, watch, nextTick } from 'vue';
  import { useI18n } from 'vue-i18n';
  import { Message } from '@arco-design/web-vue';
  import useLoading from '@/hooks/loading';
  import { ApiRecord, ApiListRes, queryApiList, saveApi,delApi } from '@/api/auth';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';

  type Column = TableColumnData & { checked?: true };
  const EditForm = ref();
  const showForm = ref(false);
  const apiListRecord = ref<ApiListRes>({});
  const { loading, setLoading } = useLoading(true);
  const { t } = useI18n();
  const formModel = ref<ApiRecord>({
    title: '',
    method: '',
    path: '',
    scope: '',
    group: '',
    need_login: 0,
    need_auth: 0,
  });
  const cloneColumns = ref<Column[]>([]);
  interface IExpandKeys {
    [key: string]: {
      [innerKey: string]: string;
    };
  }
  const expandKeys = ref<IExpandKeys>({});
  const showColumns = ref<Column[]>([]);

  const columns = computed<TableColumnData[]>(() => [
    {
      title: t('searchTable.columns.index'),
      dataIndex: 'id',
    },
    {
      title: 'api 名称',
      dataIndex: 'title',
    },
    {
      title: '请求方式',
      dataIndex: 'method',
      slotName: 'method',
    },
    {
      title: 'path',
      dataIndex: 'path',
    },
    {
      title: '登录',
      dataIndex: 'need_login',
      slotName: 'need_login',
    },
    {
      title: '鉴权',
      dataIndex: 'need_auth',
      slotName: 'need_auth',
    },
    {
      title: t('searchTable.columns.operations'),
      dataIndex: 'operations',
      fixed: 'right',
      width: 200,
      slotName: 'operation',
    },
  ]);
  const expandGroup = (scope: string, group: string) => {
    if (expandKeys.value[scope][group]) {
      delete expandKeys.value[scope][group];
    } else {
      expandKeys.value[scope][group] = 'show';
    }
  };
  const expandScope = (scope: string) => {
    if (expandKeys.value[scope]) {
      delete expandKeys.value[scope];
    } else {
      expandKeys.value[scope] = {};
    }
  };
  const handleSave = () => {
    EditForm.value.validate(async (valid: any) => {
      if (valid) {
        return;
      }
      await saveApi(formModel.value).then((res: any) => {
        if (res.code === 20000) {
          showForm.value = false;
          Message.success('保存成功');
          fetchData();
        } else {
          Message.error('修改失败');
        }
      });
    });
  };
  const del = async (record: ApiRecord) => {
    const resp=await delApi(record.id);
    if (resp.code === 20000) {
      Message.success('删除成功');
      await fetchData();
    } else {
      Message.error('删除失败');
    }
  };
  const fetchData = async () => {
    setLoading(true);
    try {
      const resp = await queryApiList();
      apiListRecord.value = resp.data;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };
  const showEdit = (record: ApiRecord) => {
    formModel.value = cloneDeep(record);
    showForm.value = true;
  };
  fetchData();

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
    name: 'ApiList',
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
</style>
