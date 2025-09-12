<template>
  <ele-card header="Official website bottom map">
    <div ref="locationMapRef" style="height: 360px; max-width: 800px"></div>
  </ele-card>
</template>

<script setup>
import { ref, watch, onMounted, onBeforeUnmount } from 'vue';
import AMapLoader from '@amap/amap-jsapi-loader';
import { useThemeStore } from '@/store/modules/theme';
import { storeToRefs } from 'pinia';
import { MAP_KEY } from '@/config/setting';

/** Is it a dark theme? */
const themeStore = useThemeStore();
const { darkMode } = storeToRefs(themeStore);

/** Map container */
const locationMapRef = ref(null);

/** Official website bottom map instance */
let mapInsLocation;

/** Render official website bottom map */
const renderLocationMap = () => {
  AMapLoader.load({
    key: MAP_KEY,
    version: '2.0',
    plugins: ['AMap.InfoWindow', 'AMap.Marker']
  })
    .then((AMap) => {
      // Render map
      const option = {
        zoom: 13, // Initial zoom level
        center: [114.346084, 30.511215 + 0.005], // Initial center point
        mapStyle: darkMode.value ? 'amap://styles/dark' : void 0
      };
      mapInsLocation = new AMap.Map(locationMapRef.value, option);
      // Create info window
      const infoWindow = new AMap.InfoWindow({
        content: ``
      });
      infoWindow.open(mapInsLocation, [114.346084, 30.511215]);
      const icon = new AMap.Icon({
        size: new AMap.Size(25, 34),
        image: '//a.amap.com/jsapi_demos/static/demo-center/icons/poi-marker-red.png',
        imageSize: new AMap.Size(25, 34)
      });
      const marker = new AMap.Marker({
        icon: icon,
        position: [114.346084, 30.511215],
        offset: new AMap.Pixel(-12, -28)
      });
      marker.setMap(mapInsLocation);
      marker.on('click', () => {
        infoWindow.open(mapInsLocation);
      });
    })
    .catch((e) => {
      console.error(e);
    });
};

/** Render map */
onMounted(() => {
  renderLocationMap();
});

/** Destroy map */
onBeforeUnmount(() => {
  if (mapInsLocation) {
    mapInsLocation.destroy();
    mapInsLocation = null;
  }
});

/** Synchronize framework dark mode switch */
watch(darkMode, (value) => {
  if (mapInsLocation) {
    if (value) {
      mapInsLocation.setMapStyle('amap://styles/dark');
    } else {
      mapInsLocation.setMapStyle('amap://styles/normal');
    }
  }
});
</script>
