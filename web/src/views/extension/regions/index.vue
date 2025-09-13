<template>
  <ele-page style="position: relative; z-index: 1">
    <ele-card header="Cascade selection of provinces, cities and districts">
      <div style="max-width: 280px">
        <regions-select v-model="city" placeholder="Please select a province, city and district" :teleported="false" />
      </div>
      <div style="margin-top: 12px">
        <el-button type="primary" @click="setCity">Display data</el-button>
      </div>
    </ele-card>
    <ele-card header="Cascade selection of provinces and cities">
      <div style="max-width: 280px">
        <regions-select v-model="provinceCity" :filterable="false" placeholder="Please select a province and city"
          type="provinceCity" :teleported="false" />
      </div>
      <div style="margin-top: 12px">
        <el-button type="primary" @click="setProvinceCity">Display data</el-button>
      </div>
    </ele-card>
    <ele-card header="Province selection">
      <div style="max-width: 280px">
        <regions-select v-model="province" :filterable="false" placeholder="Please select a province" type="province"
          :teleported="false" />
      </div>
      <div style="margin-top: 12px">
        <el-button type="primary" @click="setProvince">Display data</el-button>
      </div>
    </ele-card>
    <ele-card header="Multiple selection">
      <div style="max-width: 280px">
        <regions-select v-model="citys" :filterable="false" placeholder="Please select a province, city and district"
          :cascader-props="cascaderProps" :teleported="false" />
      </div>
      <div style="margin-top: 12px">
        <el-button type="primary" @click="setCitys">Display data</el-button>
      </div>
      <div style="margin-top: 28px">Multiple selection of provinces and cities：</div>
      <div style="max-width: 280px; margin-top: 12px">
        <regions-select v-model="provinceCitys" :filterable="false" placeholder="Please select a province and city"
          type="provinceCity" :cascader-props="cascaderProps" :teleported="false" />
      </div>
      <div style="margin-top: 28px">Multiple selection of provinces：</div>
      <div style="max-width: 280px; margin-top: 12px">
        <regions-select v-model="provinces" :filterable="false" placeholder="Please select a province" type="province"
          :cascader-props="cascaderProps" :teleported="false" />
      </div>
    </ele-card>
    <ele-card header="Translation">
      <div style="margin: 0 0 16px 0; opacity: 0.6">Single selection translation</div>
      <regions-select component="text" :model-value="['420000', '420100', '420103']" />
      <div style="margin: 32px 0 16px 0; opacity: 0.6">Multiple selection translation</div>
      <regions-select component="text" :model-value="[
        ['420000', '420100', '420103'],
        ['420000', '420100', '420106'],
        ['420000', '420100', '420111']
      ]" />
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref, reactive } from 'vue';
import RegionsSelect from '@/components/RegionsSelect/index.vue';

/** Selected province, city and district */
const city = ref([]);

/** Display data */
const setCity = () => {
  city.value = ['420000', '420100', '420103'];
};

/** Selected province and city */
const provinceCity = ref([]);

/** Display data */
const setProvinceCity = () => {
  provinceCity.value = ['420000', '420100'];
};

/** Selected province */
const province = ref([]);

/** Display data */
const setProvince = () => {
  province.value = ['420000'];
};

/** Selected province, city and district */
const citys = ref([]);

/** Display data */
const setCitys = () => {
  citys.value = [
    ['420000', '420100', '420103'],
    ['420000', '420100', '420106'],
    ['420000', '420100', '420111']
  ];
};

/** Cascader properties */
const cascaderProps = reactive({ multiple: true });

/** Selected province and city */
const provinceCitys = ref([]);

/** Selected province */
const provinces = ref([]);
</script>

<script>
export default {
  name: 'ExtensionRegions'
};
</script>
