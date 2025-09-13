<template>
  <ele-page>
    <ele-card header="Single line text ellipsis">
      <option-item label="Tooltip style">
        <el-radio-group v-model="tooltip">
          <el-radio :value="1" label="Native" />
          <el-radio :value="2" label="Tooltip" />
          <el-radio :value="3" label="Close" />
        </el-radio-group>
      </option-item>
      <div style="width: 520px; max-width: 100%; margin-top: 8px">
        <ele-ellipsis :tooltip="tooltipProps">
          The text ellipsis component `ele-ellipsis` supports single-line and multi-line ellipsis,
          when the mouse moves over it, it will prompt for all content, and supports using the Tooltip component for
          prompting.
        </ele-ellipsis>
      </div>
    </ele-card>
    <ele-card header="Multi-line text ellipsis">
      <option-item label="Maximum number of lines">
        <el-radio-group v-model="maxLine">
          <el-radio :value="2" label="2" />
          <el-radio :value="3" label="3" />
          <el-radio :value="4" label="4" />
        </el-radio-group>
      </option-item>
      <div style="width: 520px; max-width: 100%; margin-top: 8px">
        <ele-ellipsis :tooltip="tooltipProps" :max-line="maxLine">
          The text ellipsis component `ele-ellipsis` supports single-line and multi-line ellipsis,
          when the mouse moves over it, it will prompt for all content, and supports using the Tooltip component for
          prompting,
          and supports setting the other attributes of the Tooltip component, such as the position and theme,
          and can also set the content of the prompt and the content of the display to be completely different,
          and the default is to get the text content to be displayed as the prompt content.
        </ele-ellipsis>
      </div>
    </ele-card>
    <ele-card header="Text component">
      <div style="line-height: 1.8" class="demo-text-group">
        <ele-text :icon="UserOutlined">&nbsp;Regular text</ele-text>
        <ele-text :icon="UserOutlined" type="secondary">&nbsp;Secondary text</ele-text>
        <ele-text :icon="UserOutlined" type="placeholder">&nbsp;Placeholder text</ele-text>
        <ele-text :icon="UserOutlined" type="heading">&nbsp;Heading text</ele-text>
        <ele-text :icon="UserOutlined" type="primary">&nbsp;Primary text</ele-text>
        <ele-text :icon="UserOutlined" type="success">&nbsp;Success text</ele-text>
        <ele-text :icon="UserOutlined" type="warning">&nbsp;Warning text</ele-text>
        <ele-text :icon="UserOutlined" type="danger">&nbsp;Danger text</ele-text>
        <ele-text :icon="UserOutlined" type="info">&nbsp;Info text</ele-text>
        <ele-text deleted>Deleted text</ele-text>
        <ele-text underline>Underlined text</ele-text>
        <ele-text deleted underline>Deleted and underlined text</ele-text>
        <ele-text strong>Bold text</ele-text>
        <ele-text italic>Italic text</ele-text>
        <ele-text size="xs">Extra small text</ele-text>
        <ele-text size="sm">Small text</ele-text>
        <ele-text>Regular text</ele-text>
        <ele-text size="md">Medium text</ele-text>
        <ele-text size="lg">Large text</ele-text>
        <ele-text size="xl">Extra large text</ele-text>
        <ele-text size="xxl">Extra extra large text</ele-text>
      </div>
    </ele-card>
    <ele-card header="Text copy">
      <div>
        <ele-copyable>This is a copyable text.</ele-copyable>
      </div>
      <div style="margin: 16px 0">
        <ele-copyable style="width: 100%; max-width: 520px">
          <ele-ellipsis :tooltip="tooltipProps">
            The text ellipsis component `ele-ellipsis` supports single-line and multi-line ellipsis,
            when the mouse moves over it, it will prompt for all content, and supports using the Tooltip component for
            prompting.
          </ele-ellipsis>
        </ele-copyable>
      </div>
      <div style="display: flex; align-items: center; max-width: 280px">
        <el-input v-model="text" placeholder="Enter text" style="flex: 1; margin-right: 8px" />
        <ele-copyable :text="text" :custom-style="{ margin: 0 }" />
      </div>
    </ele-card>
    <ele-card header="Scrolling number">
      <ele-text size="xxl" style="padding: 0 0 16px 10px">
        <ele-count-up ref="countUpRef" :end-val="demoNum" :options="option" />
      </ele-text>
      <div>
        <el-button class="ele-btn-icon" @click="restart">Restart</el-button>
        <el-button class="ele-btn-icon" @click="update">Update number</el-button>
      </div>
    </ele-card>
    <ele-card header="Status text">
      <option-item label="Default effect">
        <el-space size="large">
          <ele-dot :ripple="ripple" text="Running" :size="dotSize" />
          <ele-dot type="success" :ripple="ripple" text="Online" :size="dotSize" />
          <ele-dot type="danger" :ripple="ripple" text="Offline" :size="dotSize" />
          <ele-dot type="info" :ripple="ripple" text="Closed" :size="dotSize" />
        </el-space>
      </option-item>
      <option-item label="Display animation">
        <el-radio-group v-model="ripple">
          <el-radio :value="true" label="Yes" />
          <el-radio :value="false" label="No" />
        </el-radio-group>
      </option-item>
      <option-item label="Size selection">
        <el-radio-group v-model="dotSize">
          <el-radio value="8px" label="8px" />
          <el-radio value="12px" label="12px" />
          <el-radio value="16px" label="16px" />
        </el-radio-group>
      </option-item>
    </ele-card>
    <ele-card header="Status dashboard">
      <div>
        <ele-dashboard type="success" style="margin-right: 18px">
          <div style="line-height: 1">
            <span style="font-size: 48px">100</span>
            <span style="font-size: 12px; margin-left: 4px">%</span>
          </div>
          <div style="margin-top: 4px">Safe</div>
        </ele-dashboard>
        <ele-dashboard type="warning" style="margin-right: 18px">
          <div style="line-height: 1">
            <span style="font-size: 48px">70</span>
            <span style="font-size: 12px; margin-left: 4px">%</span>
          </div>
          <div style="margin-top: 4px">To be optimized</div>
        </ele-dashboard>
        <ele-dashboard type="danger">
          <div style="line-height: 1">
            <span style="font-size: 48px">40</span>
            <span style="font-size: 12px; margin-left: 4px">%</span>
          </div>
          <div style="margin-top: 4px">High risk</div>
        </ele-dashboard>
      </div>
      <div style="margin: 32px 0 14px 0; max-height: 100%">
        Customize color and size:
      </div>
      <div>
        <ele-dashboard color="#722ED1" style="margin-right: 18px">
          <div style="font-size: 48px">100</div>
        </ele-dashboard>
        <ele-dashboard size="108px" space="14px">
          <div style="font-size: 38px">100</div>
        </ele-dashboard>
      </div>
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref, computed, reactive } from 'vue';
import { UserOutlined } from '@/components/icons';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';

/** 提示风格 */
const tooltip = ref(1);

/** 最大行数 */
const maxLine = ref(2);

/** tooltip属性 */
const tooltipProps = computed(() => {
  if (tooltip.value === 3) {
    return false;
  }
  return {
    original: tooltip.value === 1,
    popperStyle: {
      width: '420px',
      maxWidth: 'calc(100vw - 32px)'
    }
  };
});

/** 输入框 */
const text = ref('EleAdminPlus');

/** 滚动数字组件实例 */
const countUpRef = ref(null);

/** 滚动数字值 */
const demoNum = ref(6317);

/** 滚动数字配置 */
const option = reactive({
  useEasing: true,
  useGrouping: true,
  separator: ',',
  decimal: '.',
  prefix: '',
  suffix: ''
});

/** 更新滚动数字 */
const update = () => {
  demoNum.value += 100 + Math.round(Math.random() * 300);
};

/** 重新开始滚动数字 */
const restart = () => {
  if (countUpRef.value) {
    countUpRef.value.reset();
    countUpRef.value.start();
  }
};

/** 状态圆点是否显示动画 */
const ripple = ref(true);

/** 状态圆点尺寸 */
const dotSize = ref('12px');
</script>

<script>
export default {
  name: 'ExtensionText'
};
</script>

<style lang="scss" scoped>
.demo-text-group>div {
  max-height: 100%;
}
</style>
