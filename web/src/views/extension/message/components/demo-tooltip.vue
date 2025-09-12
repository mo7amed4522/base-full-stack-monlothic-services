<template>
  <ele-card header="Tooltip demo">
    <option-item label="Usage">
      <ele-tooltip content="EleAdminPlus" placement="top" :offset="8">
        <el-button>Dark</el-button>
      </ele-tooltip>
      <ele-tooltip content="EleAdminPlus" placement="top" effect="light" :offset="8">
        <el-button>Light</el-button>
      </ele-tooltip>
    </option-item>
    <option-item label="Theme">
      <ele-tooltip content="EleAdminPlus" placement="top" effect="danger" :offset="8">
        <el-button class="ele-btn-icon">Error</el-button>
      </ele-tooltip>
      <ele-tooltip content="EleAdminPlus" placement="top" effect="warning" :offset="8">
        <el-button class="ele-btn-icon">Warning</el-button>
      </ele-tooltip>
      <ele-tooltip content="EleAdminPlus" placement="top" effect="success" :offset="8">
        <el-button class="ele-btn-icon">Success</el-button>
      </ele-tooltip>
      <ele-tooltip content="EleAdminPlus" placement="top" bg="linear-gradient( 135deg, #43CBFF 10%, #9708CC 100%)"
        arrow-bg="#7556E0" :offset="8">
        <el-button class="ele-btn-icon">Custom</el-button>
      </ele-tooltip>
    </option-item>
    <option-item label="Popover">
      <ele-popover :width="180" trigger="click" title="Popover title" :offset="8"
        :popper-options="{ strategy: 'fixed' }">
        <template #reference>
          <el-button class="ele-btn-icon">Popover</el-button>
        </template>
        <div>This is content.</div>
        <div>This is content.</div>
      </ele-popover>
      <ele-popconfirm :width="190" :offset="8" title="Are you sure to delete this user?"
        :popper-options="{ strategy: 'fixed' }">
        <template #reference>
          <el-button class="ele-btn-icon">PopConfirm</el-button>
        </template>
      </ele-popconfirm>
      <ele-popconfirm :width="258" :offset="8" title="User satisfaction survey"
        content="Are you satisfied with this product?" cancel-button-text="Close" confirm-button-text="Satisfied"
        :cancel-button-props="{ style: { order: -1, margin: '0 auto 0 0' } }" :popper-options="{ strategy: 'fixed' }">
        <template #reference>
          <el-button class="ele-btn-icon">PopConfirm</el-button>
        </template>
        <template #action="{ cancel }">
          <el-button size="small" plain type="danger" @click="cancel()">
            Complain
          </el-button>
        </template>
      </ele-popconfirm>
    </option-item>
  </ele-card>
</template>

<script setup>
import OptionItem from '@/views/extension/avatar/components/option-item.vue';
</script>
