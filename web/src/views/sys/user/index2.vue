<template>
  <div>
    <warning-bar title="Note: The avatar in the upper right corner can be used to switch characters" />
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="Username">
          <el-input v-model="searchInfo.username" placeholder="Username" />
        </el-form-item>
        <el-form-item label="Nickname">
          <el-input v-model="searchInfo.nickname" placeholder="Nickname" />
        </el-form-item>
        <el-form-item label="Phone">
          <el-input v-model="searchInfo.phone" placeholder="Phone" />
        </el-form-item>
        <el-form-item label="Email">
          <el-input v-model="searchInfo.email" placeholder="Email" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">
            Search
          </el-button>
          <el-button icon="refresh" @click="onReset">
            Reset
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addUser">Add User</el-button>
      </div>
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="Avatar" min-width="75">
          <template #default="scope">
            <CustomPic style="margin-top:8px" :pic-src="scope.row.headerImg" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="ID" min-width="50" prop="ID" />
        <el-table-column align="left" label="Username" min-width="150" prop="userName" />
        <el-table-column align="left" label="Nickname" min-width="150" prop="nickName" />
        <el-table-column align="left" label="Phone" min-width="180" prop="phone" />
        <el-table-column align="left" label="Email" min-width="180" prop="email" />
        <el-table-column align="left" label="Role" min-width="200">
          <template #default="scope">
            <el-cascader v-model="scope.row.authorityIds" :options="authOptions" :show-all-levels="false" collapse-tags
              :props="{ multiple: true, checkStrictly: true, label: 'authorityName', value: 'authorityId', disabled: 'disabled', emitPath: false }"
              :clearable="false" @visible-change="(flag) => { changeAuthority(scope.row, flag, 0) }"
              @remove-tag="(removeAuth) => { changeAuthority(scope.row, false, removeAuth) }" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="Enable" min-width="150">
          <template #default="scope">
            <el-switch v-model="scope.row.enable" inline-prompt :active-value="1" :inactive-value="2"
              @change="() => { switchEnable(scope.row) }" />
          </template>
        </el-table-column>

        <el-table-column label="Action" min-width="250" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="delete" @click="deleteUserFunc(scope.row)">Delete</el-button>
            <el-button type="primary" link icon="edit" @click="openEdit(scope.row)">Edit</el-button>
            <el-button type="primary" link icon="magic-stick" @click="resetPasswordFunc(scope.row)">Reset
              Password</el-button>
          </template>
        </el-table-column>

      </el-table>
      <div class="gva-pagination">
        <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
          layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-drawer v-model="addUserDialog" size="60%" :show-close="false" :close-on-press-escape="false"
      :close-on-click-modal="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">User</span>
          <div>
            <el-button @click="closeAddUserDialog">Cancel</el-button>
            <el-button type="primary" @click="enterAddUserDialog">Confirm</el-button>
          </div>
        </div>
      </template>

      <el-form ref="userForm" :rules="rules" :model="userInfo" label-width="80px">
        <el-form-item v-if="dialogFlag === 'add'" label="Username" prop="userName">
          <el-input v-model="userInfo.userName" />
        </el-form-item>
        <el-form-item v-if="dialogFlag === 'add'" label="Password" prop="password">
          <el-input v-model="userInfo.password" />
        </el-form-item>
        <el-form-item label="Nickname" prop="nickName">
          <el-input v-model="userInfo.nickName" />
        </el-form-item>
        <el-form-item label="Phone" prop="phone">
          <el-input v-model="userInfo.phone" />
        </el-form-item>
        <el-form-item label="Email" prop="email">
          <el-input v-model="userInfo.email" />
        </el-form-item>
        <el-form-item label="User Role" prop="authorityId">
          <el-cascader v-model="userInfo.authorityIds" style="width:100%" :options="authOptions"
            :show-all-levels="false"
            :props="{ multiple: true, checkStrictly: true, label: 'authorityName', value: 'authorityId', disabled: 'disabled', emitPath: false }"
            :clearable="false" />
        </el-form-item>
        <el-form-item label="Enable" prop="disabled">
          <el-switch v-model="userInfo.enable" inline-prompt :active-value="1" :inactive-value="2" />
        </el-form-item>
        <el-form-item label="Avatar" label-width="80px">
          <SelectImage v-model="userInfo.headerImg" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>

import {
  getUserList,
  setUserAuthorities,
  register,
  deleteUser,
  setUserInfo,
  resetPassword,
  getAuthorityList,
} from '@/api/user'

// import CustomPic from '@/components/customPic/index.vue'
// import WarningBar from '@/components/warningBar/warningBar.vue'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
// import SelectImage from '@/components/selectImage/selectImage.vue'

defineOptions({
  name: 'User',
})

const searchInfo = ref({
  username: '',
  nickname: '',
  phone: '',
  email: ''
})

const onSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = {
    username: '',
    nickname: '',
    phone: '',
    email: ''
  }
  getTableData()
}
// 初始化相关
const setAuthorityOptions = (AuthorityData, optionsData) => {
  AuthorityData &&
    AuthorityData.forEach(item => {
      if (item.children && item.children.length) {
        const option = {
          authorityId: item.authorityId,
          authorityName: item.authorityName,
          children: []
        }
        setAuthorityOptions(item.children, option.children)
        optionsData.push(option)
      } else {
        const option = {
          authorityId: item.authorityId,
          authorityName: item.authorityName
        }
        optionsData.push(option)
      }
    })
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

watch(() => tableData.value, () => {
  setAuthorityIds()
})

const initPage = async () => {
  getTableData()
  const res = await getAuthorityList()
  setOptions(res.data)
}

initPage()

const resetPasswordFunc = (row) => {
  ElMessageBox.confirm(
    'Are you sure you want to reset the password of this user to 123456?',
    'Warning',
    {
      confirmButtonText: 'OK',
      cancelButtonText: 'Cancel',
      type: 'warning',
    }
  ).then(async () => {
    const res = await resetPassword({
      ID: row.ID,
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg,
      })
    } else {
      ElMessage({
        type: 'error',
        message: res.msg,
      })
    }
  })
}
const setAuthorityIds = () => {
  tableData.value && tableData.value.forEach((user) => {
    index.authorityIds = index.authorities && index.authorities.map(i => {
      return i.authorityId
    })
  })
}

const authOptions = ref([])
const setOptions = (authData) => {
  authOptions.value = []
  setAuthorityOptions(authData, authOptions.value)
}

const deleteUserFunc = async (row) => {
  ElMessageBox.confirm('Are you sure you want to delete this user?', 'Warning', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    type: 'warning'
  }).then(async () => {
    const res = await deleteUser({ id: row.ID })
    if (res.code === 0) {
      ElMessage.success('Delete success')
      await getTableData()
    }
  })
}

// 弹窗相关
const userInfo = ref({
  userName: '',
  password: '',
  nickName: '',
  headerImg: '',
  authorityId: '',
  authorityIds: [],
  enable: 1,
})

const rules = ref({
  userName: [
    { required: true, message: 'Please input username', trigger: 'blur' },
    { min: 5, message: 'Minimum 5 characters', trigger: 'blur' }
  ],
  password: [
    { required: true, message: 'Please input user password', trigger: 'blur' },
    { min: 6, message: 'Minimum 6 characters', trigger: 'blur' }
  ],
  nickName: [
    { required: true, message: 'Please input user nickname', trigger: 'blur' }
  ],
  phone: [
    { pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/, message: 'Please input valid phone number', trigger: 'blur' },
  ],
  email: [
    { pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g, message: 'Please input valid email', trigger: 'blur' },
  ],
  authorityId: [
    { required: true, message: 'Please select user role', trigger: 'blur' }
  ]
})
const userForm = ref(null)
const enterAddUserDialog = async () => {
  userInfo.value.authorityId = userInfo.value.authorityIds[0]
  userForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...userInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await register(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: 'Create success' })
          await getTableData()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setUserInfo(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: 'Edit success' })
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}

const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields()
  userInfo.value.headerImg = ''
  userInfo.value.authorityIds = []
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addUser = () => {
  dialogFlag.value = 'add'
  addUserDialog.value = true
}

const tempAuth = {}
const changeAuthority = async (row, flag, removeAuth) => {
  if (flag) {
    if (!removeAuth) {
      tempAuth[row.ID] = [...row.authorityIds]
    }
    return
  }
  await nextTick()
  const res = await setUserAuthorities({
    ID: row.ID,
    authorityIds: row.authorityIds
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: 'Set user role success' })
  } else {
    if (!removeAuth) {
      row.authorityIds = [...tempAuth[row.ID]]
      delete tempAuth[row.ID]
    } else {
      row.authorityIds = [removeAuth, ...row.authorityIds]
    }
  }
}

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  userInfo.value = JSON.parse(JSON.stringify(row))
  addUserDialog.value = true
}

const switchEnable = async (row) => {
  userInfo.value = JSON.parse(JSON.stringify(row))
  await nextTick()
  const req = {
    ...userInfo.value
  }
  const res = await setUserInfo(req)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: `${req.enable === 2 ? 'Disable' : 'Enable'} success` })
    await getTableData()
    userInfo.value.headerImg = ''
    userInfo.value.authorityIds = []
  }
}

</script>

<style lang="scss">
.header-img-box {
  @apply w-52 h-52 border border-solid border-gray-300 rounded-xl flex justify-center items-center cursor-pointer;
}
</style>
