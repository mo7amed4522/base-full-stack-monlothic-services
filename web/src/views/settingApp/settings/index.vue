<template>
  <ele-page>
    <div class="live-stream-container">

      <ele-card class="platform-card">
        <template #header>
          <div class="card-header">
            <div class="platform-title">
              <el-icon class="settings-icon"><Setting /></el-icon>
              <h3>{{ $t('liveStream.androidTitle') }}</h3>
            </div>
            <p>{{ $t('liveStream.androidDescription') }}</p>
          </div>
        </template>

        <div v-if="isLoading" class="loading-container">
          <el-skeleton v-for="i in 2" :key="i" animated>
            <template #template>
              <el-skeleton-item variant="rect" style="height: 150px; width: 100%; margin-bottom: 16px" />
            </template>
          </el-skeleton>
        </div>

        <div v-else-if="error" class="error-container">
          <el-alert
            :title="$t('liveStream.errorTitle')"
            :description="$t('liveStream.errorDescription')"
            type="error"
            show-icon
          />
        </div>

        <div v-else class="settings-content">
          <ele-card class="setting-card" :class="{ 'setting-enabled': androidSettings.isStreamingEnabled }">
            <div class="setting-header">
              <div class="setting-info">
                <div class="setting-icon" :class="{ 'enabled': androidSettings.isStreamingEnabled }">
                  <el-icon v-if="androidSettings.isStreamingEnabled"><VideoCamera /></el-icon>
                  <el-icon v-else><VideoCameraFilled /></el-icon>
                </div>
                <div class="setting-details">
                  <h4>{{ $t('liveStream.liveStreaming') }}</h4>
                  <p>{{ $t('liveStream.liveStreamingDescription') }}</p>
                </div>
              </div>
              <el-tag :type="androidSettings.isStreamingEnabled ? 'success' : 'info'">
                {{ androidSettings.isStreamingEnabled ? $t('liveStream.visible') : $t('liveStream.hidden') }}
              </el-tag>
            </div>
            <div class="setting-controls">
              <div class="setting-status">
                <p class="status-title">
                  {{ androidSettings.isStreamingEnabled ? $t('liveStream.broadcastingActive') : $t('liveStream.readyToStream') }}
                </p>
                <p class="status-description">
                  {{ androidSettings.isStreamingEnabled ? $t('liveStream.streamLiveDescription') : $t('liveStream.streamOfflineDescription') }}
                </p>
              </div>
              <el-switch
                v-model="androidSettings.isStreamingEnabled"
                :loading="isTogglingAndroidStreaming"
                @change="toggleAndroidSetting('streaming', $event)"
                :disabled="!hasWritePermission"
              />
            </div>
          </ele-card>
          <ele-card class="setting-card" :class="{ 'setting-enabled': androidSettings.isLiveTrackerEnabled }">
            <div class="setting-header">
              <div class="setting-info">
                <div class="setting-icon" :class="{ 'enabled': androidSettings.isLiveTrackerEnabled }">
                  <el-icon v-if="androidSettings.isLiveTrackerEnabled"><View /></el-icon>
                  <el-icon v-else><Hide /></el-icon>
                </div>
                <div class="setting-details">
                  <h4>{{ $t('liveStream.liveAnimation') }}</h4>
                  <p>{{ $t('liveStream.liveAnimationDescription') }}</p>
                </div>
              </div>
              <el-tag :type="androidSettings.isLiveTrackerEnabled ? 'success' : 'info'">
                {{ androidSettings.isLiveTrackerEnabled ? $t('liveStream.visible') : $t('liveStream.hidden') }}
              </el-tag>
            </div>
            <div class="setting-controls">
              <div class="setting-status">
                <p class="status-title">
                  {{ androidSettings.isLiveTrackerEnabled ? $t('liveStream.animationsDisplayed') : $t('liveStream.animationsHidden') }}
                </p>
                <p class="status-description">
                  {{ androidSettings.isLiveTrackerEnabled ? $t('liveStream.animationsVisibleDescription') : $t('liveStream.animationsHiddenDescription') }}
                </p>
              </div>
              <el-switch
                v-model="androidSettings.isLiveTrackerEnabled"
                :loading="isTogglingAndroidTracker"
                @change="toggleAndroidSetting('live-tracker', $event)"
                :disabled="!hasWritePermission"
              />
            </div>
          </ele-card>
          <ele-card class="setting-card countries-card" v-if="hasWritePermission">
            <template #header>
              <h4>{{ $t('liveStream.countriesTitle') }}</h4>
              <p>{{ $t('liveStream.countriesDescription') }}</p>
            </template>
            <div class="countries-selector">
              <el-select
                v-model="androidSettings.selectedCountries"
                multiple
                filterable
                :placeholder="$t('liveStream.selectCountries')"
                style="width: 100%"
                :loading="isSavingAndroidCountries"
                @change="saveAndroidCountries"
              >
                <el-option
                  v-for="country in availableCountries"
                  :key="country.code"
                  :label="`${country.name} (${country.code})`"
                  :value="country.code"
                />
              </el-select>
            </div>
          </ele-card>
        </div>
      </ele-card>
      <ele-card class="platform-card">
        <template #header>
          <div class="card-header">
            <div class="platform-title">
              <el-icon class="settings-icon"><Setting /></el-icon>
              <h3>{{ $t('liveStream.iosTitle') }}</h3>
            </div>
            <p>{{ $t('liveStream.iosDescription') }}</p>
          </div>
        </template>

        <div v-if="isLoading" class="loading-container">
          <el-skeleton v-for="i in 2" :key="i" animated>
            <template #template>
              <el-skeleton-item variant="rect" style="height: 150px; width: 100%; margin-bottom: 16px" />
            </template>
          </el-skeleton>
        </div>

        <div v-else-if="error" class="error-container">
          <el-alert
            :title="$t('liveStream.errorTitle')"
            :description="$t('liveStream.errorDescription')"
            type="error"
            show-icon
          />
        </div>
        <div v-else class="settings-content">
          <ele-card class="setting-card" :class="{ 'setting-enabled': iosSettings.isStreamingEnabled }">
            <div class="setting-header">
              <div class="setting-info">
                <div class="setting-icon" :class="{ 'enabled': iosSettings.isStreamingEnabled }">
                  <el-icon v-if="iosSettings.isStreamingEnabled"><VideoCamera /></el-icon>
                  <el-icon v-else><VideoCameraFilled /></el-icon>
                </div>
                <div class="setting-details">
                  <h4>{{ $t('liveStream.liveStreaming') }}</h4>
                  <p>{{ $t('liveStream.liveStreamingDescription') }}</p>
                </div>
              </div>
              <el-tag :type="iosSettings.isStreamingEnabled ? 'success' : 'info'">
                {{ iosSettings.isStreamingEnabled ? $t('liveStream.visible') : $t('liveStream.hidden') }}
              </el-tag>
            </div>
            <div class="setting-controls">
              <div class="setting-status">
                <p class="status-title">
                  {{ iosSettings.isStreamingEnabled ? $t('liveStream.broadcastingActive') : $t('liveStream.readyToStream') }}
                </p>
                <p class="status-description">
                  {{ iosSettings.isStreamingEnabled ? $t('liveStream.streamLiveDescription') : $t('liveStream.streamOfflineDescription') }}
                </p>
              </div>
              <el-switch
                v-model="iosSettings.isStreamingEnabled"
                :loading="isTogglingIosStreaming"
                @change="toggleIosSetting('streaming', $event)"
                :disabled="!hasWritePermission"
              />
            </div>
          </ele-card>
          <ele-card class="setting-card" :class="{ 'setting-enabled': iosSettings.isLiveTrackerEnabled }">
            <div class="setting-header">
              <div class="setting-info">
                <div class="setting-icon" :class="{ 'enabled': iosSettings.isLiveTrackerEnabled }">
                  <el-icon v-if="iosSettings.isLiveTrackerEnabled"><View /></el-icon>
                  <el-icon v-else><Hide /></el-icon>
                </div>
                <div class="setting-details">
                  <h4>{{ $t('liveStream.liveAnimation') }}</h4>
                  <p>{{ $t('liveStream.liveAnimationDescription') }}</p>
                </div>
              </div>
              <el-tag :type="iosSettings.isLiveTrackerEnabled ? 'success' : 'info'">
                {{ iosSettings.isLiveTrackerEnabled ? $t('liveStream.visible') : $t('liveStream.hidden') }}
              </el-tag>
            </div>
            <div class="setting-controls">
              <div class="setting-status">
                <p class="status-title">
                  {{ iosSettings.isLiveTrackerEnabled ? $t('liveStream.animationsDisplayed') : $t('liveStream.animationsHidden') }}
                </p>
                <p class="status-description">
                  {{ iosSettings.isLiveTrackerEnabled ? $t('liveStream.animationsVisibleDescription') : $t('liveStream.animationsHiddenDescription') }}
                </p>
              </div>
              <el-switch
                v-model="iosSettings.isLiveTrackerEnabled"
                :loading="isTogglingIosTracker"
                @change="toggleIosSetting('live-tracker', $event)"
                :disabled="!hasWritePermission"
              />
            </div>
          </ele-card>
          <ele-card class="setting-card countries-card" v-if="hasWritePermission">
            <template #header>
              <h4>{{ $t('liveStream.countriesTitle') }}</h4>
              <p>{{ $t('liveStream.countriesDescription') }}</p>
            </template>
            <div class="countries-selector">
              <el-select
                v-model="iosSettings.selectedCountries"
                multiple
                filterable
                :placeholder="$t('liveStream.selectCountries')"
                style="width: 100%"
                :loading="isSavingIosCountries"
                @change="saveIosCountries"
              >
                <el-option
                  v-for="country in availableCountries"
                  :key="country.code"
                  :label="`${country.name} (${country.code})`"
                  :value="country.code"
                />
              </el-select>
            </div>
          </ele-card>
        </div>
      </ele-card>
    </div>
  </ele-page>
</template>

<script>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import {
  Setting,
  VideoCamera,
  VideoCameraFilled,
  View,
  Hide
} from '@element-plus/icons-vue'

export default {
  name: 'LiveStreamSettings',
  components: {
    Setting,
    VideoCamera,
    VideoCameraFilled,
    View,
    Hide
  },
  setup() {
    const isLoading = ref(false)
    const error = ref(null)
    const hasWritePermission = ref(true) 
    

    const isTogglingAndroidStreaming = ref(false)
    const isTogglingAndroidTracker = ref(false)
    const isTogglingIosStreaming = ref(false)
    const isTogglingIosTracker = ref(false)
    const isSavingAndroidCountries = ref(false)
    const isSavingIosCountries = ref(false)


    const androidSettings = reactive({
      isStreamingEnabled: true,
      isLiveTrackerEnabled: false,
      selectedCountries: ['US', 'GB', 'CA']
    })

    const iosSettings = reactive({
      isStreamingEnabled: false,
      isLiveTrackerEnabled: true,
      selectedCountries: ['US', 'DE', 'FR']
    })

    const availableCountries = ref([
      { name: 'United States', code: 'US' },
      { name: 'United Kingdom', code: 'GB' },
      { name: 'Canada', code: 'CA' },
      { name: 'Australia', code: 'AU' },
      { name: 'Germany', code: 'DE' },
      { name: 'France', code: 'FR' },
      { name: 'Spain', code: 'ES' },
      { name: 'Italy', code: 'IT' },
      { name: 'Japan', code: 'JP' },
      { name: 'China', code: 'CN' },
      { name: 'Brazil', code: 'BR' },
      { name: 'Mexico', code: 'MX' },
      { name: 'India', code: 'IN' },
      { name: 'South Korea', code: 'KR' },
      { name: 'Netherlands', code: 'NL' },
      { name: 'Sweden', code: 'SE' },
      { name: 'Norway', code: 'NO' },
      { name: 'Denmark', code: 'DK' },
      { name: 'Finland', code: 'FI' },
      { name: 'Switzerland', code: 'CH' }
    ])


    const mockApiCall = async (delay = 1000) => {
      await new Promise(resolve => setTimeout(resolve, delay))

      if (Math.random() < 0.1) {
        throw new Error('API Error')
      }
      return { success: true }
    }

    const toggleAndroidSetting = async (type, value) => {
      if (!hasWritePermission.value) return
      
      try {
        if (type === 'streaming') {
          isTogglingAndroidStreaming.value = true
        } else {
          isTogglingAndroidTracker.value = true
        }
        
        await mockApiCall(800)
        
        if (type === 'streaming') {
          androidSettings.isStreamingEnabled = value
        } else {
          androidSettings.isLiveTrackerEnabled = value
        }
        
        ElMessage.success(this.$t('liveStream.successTitle'))
      } catch (error) {
        ElMessage.error(this.$t('liveStream.errorTitle'))
        if (type === 'streaming') {
          androidSettings.isStreamingEnabled = !value
        } else {
          androidSettings.isLiveTrackerEnabled = !value
        }
      } finally {
        if (type === 'streaming') {
          isTogglingAndroidStreaming.value = false
        } else {
          isTogglingAndroidTracker.value = false
        }
      }
    }

    const toggleIosSetting = async (type, value) => {
      if (!hasWritePermission.value) return
      
      try {
        if (type === 'streaming') {
          isTogglingIosStreaming.value = true
        } else {
          isTogglingIosTracker.value = true
        }
        
        await mockApiCall(800)
        
        if (type === 'streaming') {
          iosSettings.isStreamingEnabled = value
        } else {
          iosSettings.isLiveTrackerEnabled = value
        }
        
        ElMessage.success(this.$t('liveStream.successTitle'))
      } catch (error) {
        ElMessage.error(this.$t('liveStream.errorTitle'))
        if (type === 'streaming') {
          iosSettings.isStreamingEnabled = !value
        } else {
          iosSettings.isLiveTrackerEnabled = !value
        }
      } finally {
        if (type === 'streaming') {
          isTogglingIosStreaming.value = false
        } else {
          isTogglingIosTracker.value = false
        }
      }
    }

    const saveAndroidCountries = async (countries) => {
      if (!hasWritePermission.value) return
      
      try {
        isSavingAndroidCountries.value = true
        await mockApiCall(1000)
        ElMessage.success(this.$t('liveStream.countriesSaved'))
      } catch (error) {
        ElMessage.error(this.$t('liveStream.errorTitle'))
      } finally {
        isSavingAndroidCountries.value = false
      }
    }

    const saveIosCountries = async (countries) => {
      if (!hasWritePermission.value) return
      
      try {
        isSavingIosCountries.value = true
        await mockApiCall(1000)
        ElMessage.success(this.$t('liveStream.countriesSaved'))
      } catch (error) {
        ElMessage.error(this.$t('liveStream.errorTitle'))
      } finally {
        isSavingIosCountries.value = false
      }
    }

    const loadSettings = async () => {
      try {
        isLoading.value = true
        await mockApiCall(1500)
      } catch (err) {
        error.value = err
        ElMessage.error(this.$t('liveStream.errorTitle'))
      } finally {
        isLoading.value = false
      }
    }

    onMounted(() => {
      loadSettings()
    })

    return {
      isLoading,
      error,
      hasWritePermission,
      isTogglingAndroidStreaming,
      isTogglingAndroidTracker,
      isTogglingIosStreaming,
      isTogglingIosTracker,
      isSavingAndroidCountries,
      isSavingIosCountries,
      androidSettings,
      iosSettings,
      availableCountries,
      toggleAndroidSetting,
      toggleIosSetting,
      saveAndroidCountries,
      saveIosCountries
    }
  }
}
</script>

<style scoped lang="scss">
.live-stream-container {
  display: flex;
  flex-direction: column;
  gap: 24px;
  padding: 24px;
  
  @media (min-width: 1024px) {
    flex-direction: row;
    gap: 32px;
  }
}

.platform-card {
  flex: 1;
  
  .card-header {
    .platform-title {
      display: flex;
      align-items: center;
      gap: 12px;
      margin-bottom: 8px;
      
      .settings-icon {
        color: var(--el-color-primary);
        font-size: 24px;
      }
      
      h3 {
        margin: 0;
        font-size: 24px;
        font-weight: 600;
        color: #1f2937;
      }
    }
    
    p {
      margin: 0;
      color: #6b7280;
      font-size: 14px;
      line-height: 1.5;
    }
  }
}

.loading-container {
  padding: 24px;
}

.error-container {
  padding: 24px;
}

.settings-content {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 16px;
}

.setting-card {
  border: 2px solid #e5e7eb;
  transition: all 0.3s ease;
  
  &:hover {
    border-color: var(--el-color-primary-light-7);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  }
  
  &.setting-enabled {
    border-color: var(--el-color-primary-light-5);
    background: var(--el-color-primary-light-9);
  }
  
  .setting-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 16px;
    
    .setting-info {
      display: flex;
      align-items: center;
      gap: 12px;
      
      .setting-icon {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 40px;
        height: 40px;
        border-radius: 8px;
        background: #f3f4f6;
        color: #6b7280;
        transition: all 0.3s ease;
        
        &.enabled {
          background: var(--el-color-primary);
          color: white;
        }
        
        .el-icon {
          font-size: 20px;
        }
      }
      
      .setting-details {
        h4 {
          margin: 0 0 4px 0;
          font-size: 18px;
          font-weight: 600;
          color: #1f2937;
        }
        
        p {
          margin: 0;
          font-size: 14px;
          color: #6b7280;
          line-height: 1.4;
        }
      }
    }
  }
  
  .setting-controls {
    display: flex;
    align-items: center;
    justify-content: space-between;
    
    .setting-status {
      .status-title {
        margin: 0 0 4px 0;
        font-size: 14px;
        font-weight: 500;
        color: #374151;
      }
      
      .status-description {
        margin: 0;
        font-size: 12px;
        color: #6b7280;
        line-height: 1.4;
      }
    }
  }
}

.countries-card {
  .countries-selector {
    padding: 16px 0;
  }
}

@media (max-width: 768px) {
  .live-stream-container {
    padding: 16px;
    gap: 16px;
  }
  
  .platform-card {
    .card-header {
      .platform-title {
        h3 {
          font-size: 20px;
        }
      }
    }
  }
  
  .setting-card {
    .setting-header {
      flex-direction: column;
      align-items: flex-start;
      gap: 12px;
      
      .setting-info {
        width: 100%;
      }
    }
    
    .setting-controls {
      flex-direction: column;
      align-items: flex-start;
      gap: 12px;
      
      .setting-status {
        width: 100%;
      }
    }
  }
}
</style>