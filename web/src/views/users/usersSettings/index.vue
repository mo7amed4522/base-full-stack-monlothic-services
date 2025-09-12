<template>
  <ele-page>
    <search-form @search="onSearch" />
    <ele-card :flex-table="fixedHeight" :body-style="{ paddingBottom: '4px' }" :style="{
      minHeight: fixedHeight ? '380px' : void 0,
      marginBottom: fixedHeight ? '0px' : void 0
    }">
      <ele-pro-table ref="tableRef" row-key="id" :columns="columns" :datasource="datasource" :response="response"
        v-model:current="current" v-model:selections="selections" :show-overflow-tooltip="true" :border="bordered"
        :toolbar="{ theme: 'default' }" :default-sort="{ prop: 'createdAt', order: 'descending' }"
        :footer-style="{ paddingBottom: '12px' }" :tableStyle="{ paddingBottom: '0', color: 'rgba(0, 0, 0, 0.65)' }"
        :loading="loading">
        <template #userInfoSlot="{ row }">
          <div class="user-info-card">
            <div class="user-avatar">
              <el-avatar :size="48" :src="row.image">
                <template #default>
                  <span>{{ getInitials(row.name) }}</span>
                </template>
              </el-avatar>
            </div>
            <div class="user-details">
              <div class="user-name">{{ row.name }}</div>
              <div class="user-contact">{{ row.phoneNumber || row.email }}</div>
            </div>
          </div>
        </template>
        <template #providerSlot="{ row }">
          <div class="provider-badges">
            <template v-if="row.phoneNumber">
              <el-tag size="small" type="info">{{ getCountryFromPhone(row.phoneNumber) }}</el-tag>
            </template>
            <template v-else-if="row.accounts && row.accounts.length > 0">
              <el-tag v-for="(account, index) in row.accounts" :key="index" size="small" type="primary">
                {{ getProviderName(account.providerId) }}
              </el-tag>
            </template>
          </div>
        </template>
        <template #coinsSlot="{ row }">
          <span class="coins-display">
            {{ row.coins }}
          </span>
        </template>
        <template #action="{ row }">
          <el-button type="primary" size="small" @click="handleEditCoins(row)">
            {{ t('users.editCoins') }}
          </el-button>
          <el-button type="danger" size="small" @click="handleDeleteUser(row)" v-if="isDev">
            {{ t('users.delete') }}
          </el-button>
        </template>
      </ele-pro-table>
      <div v-if="!loading && users.length === 0" class="empty-state">
        <el-empty :description="t('users.noUsersFound')">
          <template #image>
            <el-icon size="64" color="#c0c4cc">
              <User />
            </el-icon>
          </template>
        </el-empty>
      </div>
    </ele-card>
    <el-dialog v-model="showEditCoinsDialog" :title="t('users.editCoinsTitle', { name: currentUser?.name })"
      width="400px">
      <el-form :model="coinsForm" label-width="80px">
        <el-form-item :label="t('users.coins')">
          <el-input-number v-model="coinsForm.coins" :min="0" :max="999999" controls-position="right"
            style="width: 100%" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showEditCoinsDialog = false">{{ t('common.cancel') }}</el-button>
        <el-button type="primary" @click="handleSaveCoins" :loading="savingCoins">
          {{ t('common.save') }}
        </el-button>
      </template>
    </el-dialog>
    <el-dialog v-model="showDeleteDialog" :title="t('users.confirmDeleteTitle')" width="400px">
      <p>{{ t('users.confirmDeleteDesc') }}</p>
      <template #footer>
        <el-button @click="showDeleteDialog = false">{{ t('common.cancel') }}</el-button>
        <el-button type="danger" @click="handleConfirmDelete" :loading="deletingUser">
          {{ t('users.delete') }}
        </el-button>
      </template>
    </el-dialog>

  </ele-page>
</template>

<script setup>
import { ref, reactive, computed, onMounted, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { EleMessage } from 'ele-admin-plus'

import { refDebounced } from '@vueuse/core'
import SearchForm from './components/search-form.vue'


const { t } = useI18n()

const tableRef = ref(null)
const bordered = ref(true)
const fixedHeight = ref(false)
const loading = ref(false)
const searchQuery = refDebounced(ref(''), 500)
const sortBy = ref('createdAt:desc')
const users = ref([])
const totalUsers = ref(0)
const current = ref(null)
const selections = ref([])


const lastWhere = reactive({})

const showEditCoinsDialog = ref(false)
const showDeleteDialog = ref(false)
const currentUser = ref(null)
const savingCoins = ref(false)
const deletingUser = ref(false)


const coinsForm = reactive({
  coins: 0
})


const isDev = import.meta.env.DEV

const response = ref({
  dataName: 'users',
  countName: 'total'
})




const getInitials = (name) => {
  if (!name) return 'U'
  return name.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2)
}

const getProviderName = (providerId) => {
  return providerId.charAt(0).toUpperCase() + providerId.slice(1)
}

const getCountryFromPhone = (phoneNumber) => {

  if (!phoneNumber) return 'Unknown'
  if (phoneNumber.startsWith('+1')) return 'US'
  if (phoneNumber.startsWith('+44')) return 'UK'
  if (phoneNumber.startsWith('+86')) return 'CN'
  return 'Other'
}

const reload = (where) => {
  tableRef.value?.reload?.({ page: 1, where })
}

const onSearch = (where) => {
  Object.assign(lastWhere, where)
  doReload()
}

const doReload = () => {
  reload(lastWhere)
}

const handleSaveCoins = async () => {
  if (!currentUser.value) return

  savingCoins.value = true
  try {
    EleMessage.success(t('users.coinsEdited'))
    showEditCoinsDialog.value = false
    reload()
  } catch (error) {
    EleMessage.error(t('users.errorEditingCoins'))
  } finally {
    savingCoins.value = false
  }
}

const handleEditCoins = (row) => {
  currentUser.value = row
  coinsForm.coins = row.coins || 0
  showEditCoinsDialog.value = true
}

const handleDeleteUser = (row) => {
  currentUser.value = row
  showDeleteDialog.value = true
}

const handleConfirmDelete = async () => {
  if (!currentUser.value) return

  deletingUser.value = true
  try {
    EleMessage.success(t('users.userDeleted'))
    showDeleteDialog.value = false
    reload()
  } catch (error) {
    EleMessage.error(t('users.errorDeletingUser'))
  } finally {
    deletingUser.value = false
  }
}




const columns = computed(() => [
  {
    type: 'selection',
    columnKey: 'selection',
    width: 50,
    align: 'center',
    fixed: 'left'
  },
  {
    columnKey: 'userInfo',
    prop: 'name',
    label: t('users.name'),
    slot: 'userInfoSlot',
    align: 'left',
    minWidth: 200
  },
  {
    columnKey: 'provider',
    prop: 'provider',
    label: t('users.provider'),
    slot: 'providerSlot',
    align: 'center',
    minWidth: 120
  },
  {
    columnKey: 'createdAt',
    prop: 'createdAt',
    label: t('users.joined'),
    align: 'center',
    minWidth: 150,
    formatter: (row) => {
      if (!row.createdAt) return 'N/A'
      return new Date(row.createdAt).toLocaleDateString()
    }
  },
  {
    columnKey: 'coins',
    prop: 'coins',
    label: t('users.coins'),
    slot: 'coinsSlot',
    align: 'center',
    minWidth: 100
  },

  {
    columnKey: 'currentXp',
    prop: 'currentXp',
    label: t('users.xp'),
    align: 'center',
    minWidth: 80,
    formatter: (row) => row.currentXp || 'N/A'
  },
  {
    columnKey: 'action',
    prop: 'action',
    label: t('users.actions'),
    slot: 'action',
    align: 'center',
    width: 200,
    fixed: 'right'
  }
])

const generateFakeUsers = () => {
  const names = [
    'Khaled Mohamed', 'Ahmed Mohamed', 'Khaled Mohamed', 'Ahmed Mohamed',
    'ALi Ali', 'Ali Ali', 'Khaled Mohamed', 'Khaled Mohamed',
    'Ahmed Mohamed', 'Ahmed Mohamed', 'Khaled Ahmed', 'Ahmed Mohamed'
  ];
  const emails = [
    'khaled@example.com', 'ahmed@example.com', 'mohamed@example.com', 'khaled@example.com',
    'mo7amed4522@example.com', 'khaled@example.com', 'ali@example.com', 'ahmed@example.com',
    'ahmed4522@example.com', 'tom@example.com', 'mohamed@example.com', 'mohamed@example.com'
  ];

  return names.map((name, index) => {
    const coins = Math.floor(Math.random() * 500) + 50;
    return {
      id: index + 1,
      name,
      email: emails[index],
      image: `https://api.dicebear.com/7.x/avataaars/svg?seed=${name.replace(' ', '')}`,
      phoneNumber: index % 3 === 0 ? `+1234567${String(index).padStart(3, '0')}` : null,
      coins,
      currentLevel: Math.floor(Math.random() * 10) + 1,
      currentXp: Math.floor(Math.random() * 3000) + 100,
      createdAt: new Date(Date.now() - Math.random() * 30 * 24 * 60 * 60 * 1000).toISOString(),
      accounts: [{ providerId: ['google', 'facebook', 'apple'][Math.floor(Math.random() * 3)] }]
    };
  });
};

const datasource = async ({ }) => {
  loading.value = true
  try {
    const mockUsers = generateFakeUsers()

    let filteredUsers = mockUsers
    if (searchQuery.value) {
      filteredUsers = mockUsers.filter(user =>
        user.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
        (user.email && user.email.toLowerCase().includes(searchQuery.value.toLowerCase()))
      )
    }

    const [sortField, sortOrder] = sortBy.value.split(':')
    filteredUsers.sort((a, b) => {
      let aVal = a[sortField]
      let bVal = b[sortField]

      if (sortField === 'createdAt') {
        aVal = new Date(aVal)
        bVal = new Date(bVal)
      }

      if (sortOrder === 'desc') {
        return bVal > aVal ? 1 : -1
      } else {
        return aVal > bVal ? 1 : -1
      }
    })

    users.value = filteredUsers
    totalUsers.value = filteredUsers.length

    return {
      users: filteredUsers,
      total: filteredUsers.length
    }
  } catch (error) {
    EleMessage.error(t('users.errorLoading'))
    return { users: [], total: 0 }
  } finally {
    loading.value = false
  }
}




watch(searchQuery, () => {
  doReload()
})


onMounted(() => {
  doReload()
})
</script>

<style lang="scss" scoped>
.page-header {
  margin-bottom: 16px;

  .search-section {
    display: flex;
    gap: 12px;
    margin-bottom: 8px;

    .search-input {
      width: 300px;
    }
  }

  .results-info {
    color: #909399;
    font-size: 14px;
  }
}

.user-info-card {
  display: flex;
  align-items: center;
  gap: 12px;

  .user-avatar {
    flex-shrink: 0;
  }

  .user-details {
    flex: 1;
    min-width: 0;

    .user-name {
      font-weight: 600;
      color: #303133;
      margin-bottom: 4px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }

    .user-contact {
      font-size: 12px;
      color: #909399;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
  }
}

.provider-badges {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
  justify-content: center;
}

.coins-display {
  font-weight: 600;
  color: #67c23a;
}

.empty-state {
  padding: 40px 0;
  text-align: center;
}

.text-muted {
  color: #909399;
}
</style>
