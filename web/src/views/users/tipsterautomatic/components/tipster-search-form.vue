<template>
    <ele-card :body-style="{ paddingBottom: '0' }">
        <el-form ref="formRef" :model="form" class="ele-form-search" @keyup.enter="search" @submit.prevent="">
            <el-row :gutter="15">
                <el-col :lg="6" :md="12">
                    <el-form-item label="Automatic Tipster Name:">
                        <el-input v-model="form.name" placeholder="Please enter automatic tipster name" clearable />
                    </el-form-item>
                </el-col>
                <el-col :lg="6" :md="12">
                    <el-form-item label="Status:">
                        <el-select v-model="form.status" placeholder="Select status" clearable style="width: 100%">
                            <el-option label="Active" value="active" />
                            <el-option label="Inactive" value="inactive" />
                        </el-select>
                    </el-form-item>
                </el-col>
                <el-col :lg="6" :md="12">
                    <el-form-item label="Min Accuracy:">
                        <el-input-number v-model="form.minAccuracy" placeholder="Min accuracy %" :min="0" :max="100"
                            style="width: 100%" />
                    </el-form-item>
                </el-col>
                <el-col :lg="6" :md="12">
                    <el-form-item class="ele-form-actions">
                        <el-button type="primary" @click="search">
                            <template #icon>
                                <ele-icon :size="14">
                                    <search-outlined />
                                </ele-icon>
                            </template>
                            Search
                        </el-button>
                        <el-button @click="reset">
                            <template #icon>
                                <ele-icon :size="14">
                                    <reload-outlined />
                                </ele-icon>
                            </template>
                            Reset
                        </el-button>
                    </el-form-item>
                </el-col>
            </el-row>
        </el-form>
    </ele-card>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { SearchOutlined, ReloadOutlined } from '@/components/icons'

const emit = defineEmits(['search'])

const formRef = ref(null)

const form = reactive({
    name: '',
    status: '',
    minAccuracy: null
})

const search = () => {
    emit('search', { ...form })
}

const reset = () => {
    formRef.value?.resetFields()
    search()
}
</script>

<style lang="scss" scoped>
.ele-form-search {
    :deep(.el-form-item) {
        margin-bottom: 18px;
    }
}
</style>