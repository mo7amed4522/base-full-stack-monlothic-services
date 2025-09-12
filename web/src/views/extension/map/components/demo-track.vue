<template>
  <ele-card header="Track display and playback">
    <div ref="trackMapRef" style="height: 360px; max-width: 800px; margin-bottom: 16px">
    </div>
    <div>
      <el-button type="primary" class="ele-btn-icon" @click="startTrackAnim">
        Start moving
      </el-button>
      <el-button type="primary" class="ele-btn-icon" @click="pauseTrackAnim">
        Pause moving
      </el-button>
      <el-button type="primary" class="ele-btn-icon" @click="resumeTrackAnim">
        Resume moving
      </el-button>
    </div>
  </ele-card>
</template>

<script setup>
import { ref, watch, onMounted, onBeforeUnmount } from 'vue';
import AMapLoader from '@amap/amap-jsapi-loader';
import { useThemeStore } from '@/store/modules/theme';
import { storeToRefs } from 'pinia';
import { MAP_KEY } from '@/config/setting';

/** Whether it is a dark theme */
const themeStore = useThemeStore();
const { darkMode } = storeToRefs(themeStore);

/** Map container */
const trackMapRef = ref(null);

/** Instance of the car track map */
let mapInsTrack;

/** Car marker */
let carMarker;

/** Track route */
const lineData = [
  [116.478935, 39.997761],
  [116.478939, 39.997825],
  [116.478912, 39.998549],
  [116.478912, 39.998549],
  [116.478998, 39.998555],
  [116.478998, 39.998555],
  [116.479282, 39.99856],
  [116.479658, 39.998528],
  [116.480151, 39.998453],
  [116.480784, 39.998302],
  [116.480784, 39.998302],
  [116.481149, 39.998184],
  [116.481573, 39.997997],
  [116.481863, 39.997846],
  [116.482072, 39.997718],
  [116.482362, 39.997718],
  [116.483633, 39.998935],
  [116.48367, 39.998968],
  [116.484648, 39.999861]
];

/** Render track playback map */
const renderTrackMap = () => {
  AMapLoader.load({
    key: MAP_KEY,
    version: '2.0',
    plugins: ['AMap.MoveAnimation', 'AMap.Marker', 'AMap.Polyline']
  })
    .then((AMap) => {
      // Render map
      const option = {
        zoom: 17,
        center: [116.478935, 39.997761],
        mapStyle: darkMode.value ? 'amap://styles/dark' : void 0
      };
      mapInsTrack = new AMap.Map(trackMapRef.value, option);
      // Create car marker
      carMarker = new AMap.Marker({
        map: mapInsTrack,
        position: [116.478935, 39.997761],
        icon: 'https://a.amap.com/jsapi_demos/static/demo-center-v2/car.png',
        offset: new AMap.Pixel(-13, -26)
      });
      // Draw track
      new AMap.Polyline({
        map: mapInsTrack,
        path: lineData,
        showDir: true,
        strokeColor: '#2288FF', // Line Color
        strokeOpacity: 1, // Line transparency
        strokeWeight: 6 // Line width
        //strokeStyle: 'solid'  // Line style
      });
      // Passed track
      const passedPolyline = new AMap.Polyline({
        map: mapInsTrack,
        showDir: true,
        strokeColor: '#44BB55', // Line Color
        strokeOpacity: 1, // Line transparency
        strokeWeight: 6 // Line width
      });
      // Car moving callback
      carMarker.on('moving', (e) => {
        passedPolyline.setPath(e.passedPath);
      });
      // Map adaptive
      mapInsTrack.setFitView();
    })
    .catch((e) => {
      console.error(e);
    });
};

/** Start track playback animation */
const startTrackAnim = () => {
  if (carMarker) {
    carMarker.stopMove();
    carMarker.moveAlong(lineData, {
      duration: 200,
      autoRotation: true
    });
  }
};

/** Pause track playback animation */
const pauseTrackAnim = () => {
  if (carMarker) {
    carMarker.pauseMove();
  }
};

/** Resume track playback animation */
const resumeTrackAnim = () => {
  if (carMarker) {
    carMarker.resumeMove();
  }
};

/** Render track playback map */
onMounted(() => {
  renderTrackMap();
});

/** Destroy track playback map */
onBeforeUnmount(() => {
  if (mapInsTrack) {
    mapInsTrack.destroy();
    mapInsTrack = null;
  }
});

/** Synchronize framework dark mode switch */
watch(darkMode, () => {
  if (mapInsTrack) {
    if (darkMode.value) {
      mapInsTrack.setMapStyle('amap://styles/dark');
    } else {
      mapInsTrack.setMapStyle('amap://styles/normal');
    }
  }
});
</script>
