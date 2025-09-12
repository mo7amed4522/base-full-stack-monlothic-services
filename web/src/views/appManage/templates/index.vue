<template>
  <ele-page>
    <!-- Search form -->
    <search-form @search="onSearch" />
    <ele-card :flex-table="fixedHeight" :body-style="{ paddingBottom: '4px' }" :style="{
      minHeight: fixedHeight ? '380px' : void 0,
      marginBottom: fixedHeight ? '0px' : void 0
    }">
      <ele-pro-table ref="tableRef" row-key="id" :columns="columns" :datasource="datasource" :response="response"
        v-model:current="current" v-model:selections="selections" :show-overflow-tooltip="true" :border="bordered"
        :toolbar="{ theme: 'default' }" :default-sort="{ prop: 'createTime', order: 'descending' }"
        :footer-style="{ paddingBottom: '12px' }" :tableStyle="{ paddingBottom: '0', color: 'rgba(0, 0, 0, 0.65)' }">
        <template #toolbar>
          <el-button type="primary" :icon="PlusOutlined" class="ele-btn-icon" @click="handleEdit()">Add
            template</el-button>
        </template>
        <!-- Header tool buttons -->
        <template #tools></template>

        <template #action="{ row }">
          <el-link type="primary" :underline="false" icon="edit" @click="handleEdit(row)">Edit</el-link>
          <el-divider direction="vertical" />
          <el-link type="danger" :underline="false" icon="Delete" @click="handleDelete(row)">{{ t('common.deleted')
          }}</el-link>
        </template>
      </ele-pro-table>
    </ele-card>

    <!-- Edit form -->
    <edit-form v-model="showEdit" :data="current" @done="reload" />
  </ele-page>
</template>
<script setup>
import { ref, reactive, computed, onMounted, watch } from 'vue';
import SearchForm from './components/search-form.vue';
import { EleMessage } from 'ele-admin-plus';
import EditForm from './components/edit-form.vue';
import { getCPageTemplateListApi, deleteCPageTemplateApi } from '@/api/app'
import { useMessageBox } from 'ele-admin-plus/es';
import { useI18n } from 'vue-i18n';
import { _timeFormat } from '@/utils/public'
import { PlusOutlined } from '@/components/icons';

const { t } = useI18n();

const messageBox = useMessageBox();
/** Table instance */
const tableRef = ref(null);
const bordered = ref(true);

/** Table search parameters */
const lastWhere = reactive({});
/** Table selected data */
const selections = ref([]);
/** Table single selected data */
const current = ref(null);
/** Table fixed height */
const fixedHeight = ref(false);
/** Refresh table */
const reload = (where) => {
  tableRef.value?.reload?.({ page: 1, where });
};

/** Search event */
const onSearch = (where) => {
  Object.assign(lastWhere, where);
  doReload();
};

/** Table search */
const doReload = () => {
  reload(lastWhere);
};

const response = ref({
  dataName: 'list', // Field name of the data list, nested is supported, for example: 'data.rows'
  countName: 'total' // The field name of the total data, for example: 'data.total'
});

const showEdit = ref(false);

/** Open edit form */
const handleEdit = (row) => {
  current.value = row ?? null;
  showEdit.value = true;
};

/** Delete event */
const handleDelete = async (row) => {
  messageBox
    .confirm(t('common.deleteConfirmText'), t('common.deleteConfirm'), {
      type: 'warning',
      draggable: true
    })
    .then(async () => {
      try {
        const res = await deleteCPageTemplateApi({ id: row.id });
        EleMessage.success(res.msg);
        doReload();
      } catch (error) {
        EleMessage.error(error.message);
      }
    })
    .catch(() => { });
};

/** Table column configuration */
const columns = computed(() => {
  return [
    {
      type: 'selection',
      columnKey: 'selection',
      width: 56,
      align: 'center'
    },
    {
      columnKey: 'id',
      prop: 'id',
      label: 'ID',
      align: 'center',
      minWidth: 80
    },
    {
      columnKey: 'name',
      prop: 'name',
      label: 'Template name',
      align: 'center',
      minWidth: 120
    },
    {
      columnKey: 'code',
      prop: 'code',
      label: 'Template code',
      align: 'center',
      minWidth: 120
    },
    {
      columnKey: 'createdAt',
      prop: 'createdAt',
      label: t('common.createTime'),
      align: 'center',
      minWidth: 180,
      formatter: (row, column, cellValue, index) => {
        return cellValue ? _timeFormat(cellValue) : '-';
      }
    },
    {
      columnKey: 'action',
      prop: 'action',
      label: t('common.action'),
      align: 'center',
      slot: 'action',
      width: 260,
      fixed: 'right'
    }
  ];
});

/** Table data source */
const datasource = async ({ pages, where, orders, filters }) => {
  try {
    const data = await getCPageTemplateListApi({
      ...where,
      page: pages.page || 1, // Use pages.currentPage or the default value of 1
      pageSize: pages.limit || 10 // Use pages.pageSize or the default value of 10
    });
    return data;
  } catch (e) {
    EleMessage.error(e.message);
  }
};

onMounted(() => { });
</script>

<style lang="scss" scoped></style>
