<template>
  <ele-card header="Map location selection">
    <div>
      <ele-dropdown v-model="mode" :items="[
        { title: 'Click map mode', command: 'lnglat' },
        { title: 'Keyword retrieval mode', command: 'keyword' },
        { title: 'POI retrieval mode', command: 'poi' }
      ]" trigger="click" style="margin-right: 12px">
        <template #default="{ selected }">
          <el-button class="ele-btn-icon" style="width: 134px">
            <span>{{ selected.title }}</span>
            <el-icon :size="12" style="margin: 0 -4px 0 2px">
              <ArrowDown />
            </el-icon>
          </el-button>
        </template>
      </ele-dropdown>
      <el-button type="primary" class="ele-btn-icon" @click="openMapPicker">
        Open location selection
      </el-button>
    </div>
    <div style="margin-top: 12px">Longitude and latitude: {{ result.lngAndLat }}</div>
    <div style="margin-top: 12px">Selected location: {{ result.location }}</div>
    <div style="margin-top: 12px">Detailed address: {{ result.address }}</div>
  </ele-card>
  <!-- Map location selection popup -->
  <ele-map-picker v-model="visible" :return-regions="true" :selected="selectedData" :dark-mode="darkMode" :mode="mode"
    height="100%" :modal-props="{
      bodyStyle: { height: '460px', minHeight: '100%', maxHeight: '100%' },
      maxable: true,
      resizable: true,
      resizeIconStyle: { zIndex: 200 },
      minWidth: 372,
      minHeight: 280
    }" :side-style="{ maxWidth: '448px' }" @done="onChoose" />
</template>

<script setup>
import { ref, reactive } from 'vue';
import { ArrowDown } from '@/components/icons';
import { useThemeStore } from '@/store/modules/theme';
import { storeToRefs } from 'pinia';

const themeStore = useThemeStore();
const { darkMode } = storeToRefs(themeStore);

/** Whether to display the map selection popup */
const visible = ref(false);

/** Location retrieval type */
const mode = ref('lnglat');

/** Recipient of the selected location */
const selectedData = ref();

/** Selection result */
const result = reactive({
  location: '',
  address: '',
  lngAndLat: ''
});

/** Open location selection */
const openMapPicker = () => {
  visible.value = true;
};

/** Map selection callback */
const onChoose = (location) => {
  console.log(JSON.parse(JSON.stringify(location)));
  result.location = `${location.city?.province}/${location.city?.city}/${location.city?.district}`;
  result.address = `${location.name} ${location.address}`;
  result.lngAndLat = `${location.lng},${location.lat}`;
  visible.value = false;
  selectedData.value = location;
};
</script>
