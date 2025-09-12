<template>
  <ele-page>
    <ele-card header="Basic Usage">
      <ele-check-card multiple v-model="checked" :items="items" :row="{ gutter: 12 }">
        <template #item="{ item }">
          <div style="display: flex; padding: 18px 12px; height: 87px">
            <el-avatar :size="46" shape="square" :src="item.avatar" />
            <div style="flex: 1; overflow: hidden; padding-left: 12px">
              <ele-text size="md">{{ item.title }}</ele-text>
              <ele-ellipsis size="sm" type="secondary" :max-line="3">
                {{ item.content }}
              </ele-ellipsis>
            </div>
          </div>
        </template>
      </ele-check-card>
      <div>Selected: {{ JSON.stringify(checked) }}</div>
    </ele-card>
    <ele-card header="Single Selection Mode" :collapsable="true" :body-style="{ padding: 0 }">
      <div style="padding: 20px">
        <ele-check-card v-model="checked2" :items="items" :row="{ gutter: 12 }">
          <template #item="{ item }">
            <div style="display: flex; padding: 18px 12px; height: 87px">
              <el-avatar :size="46" shape="square" :src="item.avatar" />
              <div style="flex: 1; overflow: hidden; padding-left: 12px">
                <ele-text size="md">{{ item.title }}</ele-text>
                <ele-ellipsis size="sm" type="secondary" :max-line="3">
                  {{ item.content }}
                </ele-ellipsis>
              </div>
            </div>
          </template>
        </ele-check-card>
        <div>Selected: {{ checked2 }}</div>
      </div>
    </ele-card>
    <ele-card header="Disabled" collapsable="header" :header-style="{ cursor: 'pointer', userSelect: 'none' }"
      :body-style="{ padding: 0 }">
      <div style="padding: 20px">
        <ele-check-card :items="items2" v-model="checked2" :disabled="disabled" :row="{ gutter: 12 }">
          <template #item="{ item }">
            <div style="display: flex; padding: 18px 12px; height: 87px">
              <el-avatar :size="46" shape="square" :src="item.avatar" />
              <div style="flex: 1; overflow: hidden; padding-left: 12px">
                <ele-text size="md">{{ item.title }}</ele-text>
                <ele-ellipsis size="sm" type="secondary" :max-line="3">
                  {{ item.content }}
                </ele-ellipsis>
              </div>
            </div>
          </template>
        </ele-check-card>
        <div style="display: flex; align-items: center">
          <div>Overall disabled:</div>
          <el-switch v-model="disabled" size="small" />
        </div>
      </div>
    </ele-card>
    <ele-card header="Tag Style" collapsable="header" :header-style="{ cursor: 'pointer', userSelect: 'none' }"
      :body-style="{ padding: 0 }" :collapse-icon-style="{ order: -1, margin: '1px 6px 0 -10px' }">
      <div style="padding: 28px 20px">
        <div style="display: flex">
          <div style="flex-shrink: 0">Personal hobbies:</div>
          <ele-check-card :arrow="false" :bordered="false" v-model="checked3" :items="items4" class="check-tag" />
        </div>
        <div style="display: flex; margin-top: 12px">
          <div style="flex-shrink: 0">Multiple selection mode:</div>
          <ele-check-card multiple :arrow="false" :bordered="false" v-model="checked4" :items="items4"
            class="check-tag" />
        </div>
      </div>
      <template #footer>
        <div>Footer</div>
      </template>
    </ele-card>
    <ele-card header="Custom Content">
      <div style="display: flex">
        <div style="flex-shrink: 0; margin-top: 20px">Payment method:&emsp;</div>
        <ele-check-card v-model="checked5" :items="items5"
          :item-style="{ margin: '8px', padding: '8px 12px', display: 'flex' }" style="display: flex; flex-wrap: wrap">
          <template #item="{ item }">
            <div :style="{
              width: '98px',
              height: '28px',
              background: `center / contain no-repeat url(${item.img})`
            }" class="demo-card-pay-img"></div>
          </template>
        </ele-check-card>
      </div>
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref } from 'vue';
import { getDemoCardData } from '@/api/example';

/** Card selection */
const checked = ref([2]);

/** Card data */
const items = ref([]);

// Get card data
getDemoCardData()
  .then((data) => {
    items.value = data.map((d) => {
      return {
        ...d,
        col: { md: 6, sm: 12, xs: 24, style: { marginBottom: '12px' } }
      };
    });
    items2.value = items.value.map((d) => {
      return { ...d, disabled: d.value === 2 };
    });
  })
  .catch((e) => {
    console.error(e);
  });

/** Single selection */
const checked2 = ref();

/** Disabled data */
const items2 = ref([]);

/** Overall disabled */
const disabled = ref(false);

/** Tag style single selection */
const checked3 = ref('See a movie');

/** Tag style multiple selection */
const checked4 = ref(['See a movie']);

/** Personal hobbies data */
const items4 = ref([
  { value: 'See a movie' },
  { value: 'Listen to music' },
  { value: 'Play basketball' },
  { value: 'Read' }
]);

/** Payment method */
const checked5 = ref();

/** Payment method data */
const items5 = ref([
  {
    value: 1,
    img: 'https://gw.alipayobjects.com/mdn/member_frontWeb/afts/img/A*oRlnSYAsgYQAAAAAAAAAAABkARQnAQ'
  },
  {
    value: 2,
    img: 'https://e.gitee.com/assets/images/wechatpay.png'
  },
  {
    value: 3,
    img: 'https://cn.unionpay.com/upowhtml/cn/resources/images/header/homepage-logo.png'
  }
]);
</script>

<script>
export default {
  name: 'ExtensionCheckCard'
};
</script>

<style lang="scss" scoped>
/* Tag style */
.check-tag {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  box-sizing: border-box;
  padding-left: 12px;

  :deep(.ele-check-card) {
    font-size: 13px;
    height: 23px;
    line-height: 23px;
    padding: 0 8px;
    margin: 0 8px 8px 0;
    border-radius: 4px;
    transition: (color 0.2s, background-color 0.2s);

    &:not(.is-disabled):not(.is-checked):hover {
      background: var(--el-fill-color-light);
    }

    &.is-checked {
      color: #fff;
      background: var(--el-color-primary);
    }

    &.is-disabled {
      background: none;
      opacity: 1;

      &.is-checked {
        color: var(--el-disabled-text-color);
        background: var(--el-disabled-bg-color);
      }
    }
  }
}
</style>

<style lang="scss">
/* Payment method images adapted to dark theme */
html.dark .demo-card-pay-img {
  display: flex;
  align-items: center;
  justify-content: center;
  background-position-y: -28px !important;

  &::before,
  &::after {
    content: '';
    width: 40px;
    height: 100%;
    background: inherit;
    background-size: auto 28px;
    background-position: 9px 0;
    flex-shrink: 0;
    overflow: hidden;
  }

  &::after {
    flex: 1;
    width: 100%;
    background-position: -30px 0;
    filter: invert(1);
  }
}

.ele-check-card-group>.ele-check-card:nth-child(2) .demo-card-pay-img {
  &::before {
    width: 32px;
    background-size: auto 26px;
    background-position: 0px 1px;
  }

  &::after {
    background-size: auto 26px;
    background-position: -32px 0;
  }
}

.ele-check-card-group>.ele-check-card:last-child .demo-card-pay-img {
  &::before {
    width: 38px;
    background-size: auto 24px;
    background-position: 0px 2px;
  }

  &::after {
    background-size: auto 24px;
    background-position: -36px 1px;
  }
}
</style>
