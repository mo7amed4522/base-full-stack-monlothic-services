<template>
  <ele-page>
    <!-- Search form -->
    <search-form @search="onSearch" />
    <ele-card
      :flex-table="fixedHeight"
      :body-style="{ paddingBottom: '4px' }"
      :style="{
        minHeight: fixedHeight ? '380px' : void 0,
        marginBottom: fixedHeight ? '0px' : void 0
      }"
    >
      <ele-pro-table
        ref="tableRef"
        row-key="id"
        :columns="columns"
        :datasource="datasource"
        :response="response"
        v-model:current="current"
        v-model:selections="selections"
        :show-overflow-tooltip="true"
        :border="bordered"
        :toolbar="{ theme: 'default' }"
        :default-sort="{ prop: 'createTime', order: 'descending' }"
        :footer-style="{ paddingBottom: '12px' }"
        :tableStyle="{ paddingBottom: '0', color: 'rgba(0, 0, 0, 0.65)' }"
      >
        <template #toolbar>
          <el-button type="primary" :icon="PlusOutlined" class="ele-btn-icon" @click="handleEdit()">{{ t('common.add') }}</el-button>
        </template>
        <!-- Table header tool buttons -->
        <template #tools></template>

        <template #visibleSlot="{ row }">
          <ele-dot v-if="row.visible === 1" :text="t('common.enable')" type="success" size="8px" />
          <ele-dot v-else-if="row.visible === -1" :text="t('common.disable')" size="8px" type="danger" :ripple="false" />
        </template>

        <template #statusSlot="{ row }">
          <ele-dot v-if="row.status === 1" :text="t('common.match')" type="success" size="8px" />
          <ele-dot v-else-if="row.status === 2" :text="t('common.matchSuccess')" size="8px" type="success" :ripple="false" />
          <ele-dot v-else-if="row.status === 3" :text="t('common.matchFail')" size="8px" type="danger" :ripple="false" />
        </template>

        <template #action="{ row }">
          <el-link type="primary" :underline="false" icon="Plus" @click="handleAddSubDomain(row, false)">{{ t('common.addSubdomain') }}</el-link>
          <el-divider direction="vertical" />
          <el-link :type="row.visible === 1 ? 'danger' : 'primary'" icon="Open" :underline="false" @click="handleVisible(row)">
            {{ row.visible === 1 ? t('common.disable') : t('common.enable') }}
          </el-link>
          <el-divider direction="vertical" />
          <el-link type="primary" :underline="false" icon="view" @click="handleAddSubDomain(row, true)">{{ t('common.viewSettings') }}</el-link>
        </template>
      </ele-pro-table>
    </ele-card>

    <!-- Edit Domain -->
    <edit-from v-model="showEdit" :data="current" @done="reload" />

    <!-- Add a subdomain -->
    <sub-domain-form v-model="showSubDomain" :data="current" :is-view="isView" @done="reload" />
  </ele-page>
</template>
<script setup>
  import { getCDominManageListApi, visibleCDominManageApi } from '@/api/app/index.js';
  import { PlusOutlined } from '@/components/icons';
  import { _timeFormat } from '@/utils/public.js';
  import { EleMessage } from 'ele-admin-plus';
  import { useMessageBox } from 'ele-admin-plus/es';
  import { computed, onMounted, reactive, ref } from 'vue';
  import { useI18n } from 'vue-i18n';
  import EditFrom from './components/edit-form.vue';
  import SearchForm from './components/search-form.vue';
  import SubDomainForm from './components/subdomain-form.vue';

  const { t } = useI18n();

  const messageBox = useMessageBox();

  // Define the base URL of the image
  const imageBaseUrl = import.meta.env.VITE_APP_IMAGE_URL || '';

  // How to get the full image URL
  const getImageUrl = (imagePath) => {
    if (!imagePath) return '';
    if (imagePath.startsWith('http')) return imagePath; // If it is already a complete URL
    return imageBaseUrl + imagePath;
  };
  /** Table instance */
  const tableRef = ref(null);
  const bordered = ref(true);

  /** Table search parameters */
  const lastWhere = reactive({});
  /** Table selected data */
  const selections = ref([]);
  /** Table single selection selected data */
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
    dataName: 'list', // The field name of the data list, support nesting, such as: 'data.rows'
    countName: 'total' // The field name of the data total count, such as: 'data.total'
  });

  const showSubDomain = ref(false);
  const isView = ref(false);

  const handleAddSubDomain = (row, isViewBool) => {
    current.value = row ?? null;
    showSubDomain.value = true;
    isView.value = isViewBool;
  };

  const showEdit = ref(false);

  /** Open edit popup */
  const handleEdit = (row) => {
    current.value = row ?? null;
    showEdit.value = true;
  };

  // Show/hide using messageBox prompt
  const handleVisible = async (row) => {
    messageBox
      .confirm(row.visible === 1 ? 'Are you sure you want to hide? ' : 'Are you sure you want to show? ', 'Prompt', {
        type: 'warning',
        draggable: true
      })
      .then(async () => {
        try {
          const res = await visibleCDominManageApi({ id: row.id, visible: row.visible === 1 ? -1 : 1 });
          EleMessage.success(res.msg);
          doReload();
        } catch (error) {
          EleMessage.error(error.message);
        }
      });
  };

  // Delete
  const handleDelete = async (row) => {
    messageBox
      .confirm(t('common.deleteConfirmText'), t('common.deleteConfirm'), {
        type: 'warning',
        draggable: true
      })
      .then(async () => {
        try {
          const res = await deleteAppApi({ id: row.id });
          EleMessage.success(res.msg);
          doReload();
        } catch (error) {
          EleMessage.error(error.message);
        }
      })
      .catch(() => {});
  };

  /** Table column configuration */
  const columns = computed(() => {
    return [
      {
        columnKey: 'id',
        prop: 'id',
        label: 'ID',
        align: 'center'
      },
      {
        columnKey: 'domain',
        prop: 'domain',
        label: 'Domain name',
        align: 'center',
        minWidth: 120
      },
      {
        columnKey: 'resolveIp',
        prop: 'resolveIp',
        label: 'Resolved IP',
        align: 'center',
        minWidth: 120
      },
      {
        columnKey: 'accessType',
        prop: 'accessType',
        label: 'Access type',
        align: 'center',
        minWidth: 120,
        formatter: (row, column, cellValue, index) => {
          switch (cellValue) {
            case 1:
              return 'Domain hosting';
            default:
              return '-';
          }
        }
      },
      {
        columnKey: 'status',
        prop: 'status',
        label: 'Status',
        slot: 'statusSlot',
        align: 'center',
        minWidth: 120,
        formatter: (row, column, cellValue, index) => {
          switch (cellValue) {
            case 1:
              return 'Matching';
            case 2:
              return 'Correct match';
            case 3:
              return 'Match failed';
            default:
              return '-';
          }
        }
      },
      {
        columnKey: 'visible',
        prop: 'visible',
        label: 'Visible status',
        slot: 'visibleSlot',
        align: 'center',
        minWidth: 120,
        formatter: (row, column, cellValue, index) => {
          switch (cellValue) {
            case 1:
              return 'Visible';
            case -1:
              return 'Hidden';
            default:
              return '-';
          }
        }
      },
      {
        columnKey: 'remark',
        prop: 'remark',
        label: 'Remark',
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
        label: 'Action',
        align: 'center',
        slot: 'action',
        width: 320,
        fixed: 'right'
      }
    ];
  });

  /** Table data source */
  const datasource = async ({ pages, where, orders, filters }) => {
    try {
      const data = await getCDominManageListApi({
        ...where,
        page: pages.page || 1, // Use pages.currentPage or the default value of 1
        pageSize: pages.limit || 10 // Use pages.pageSize or the default value of 10
      });
      return data;
    } catch (e) {
      EleMessage.error(e.message);
    }
  };

  onMounted(() => {});
</script>
