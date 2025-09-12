<template>
  <ele-page>
    <div class="otp-settings-container">
      <div v-if="isLoading" class="loading-container">
        <el-skeleton :rows="6" animated />
      </div>
      <div v-else-if="error" class="error-container">
        <el-alert
          :title="$t('common.error')"
          :description="error"
          type="error"
          show-icon
          :closable="false"
        />
      </div>
      <div v-else class="otp-content">
        <ele-card class="otp-card">
          <template #header>
            <div class="card-header">
              <div class="otp-title">
                <el-icon class="otp-icon"><Message /></el-icon>
                <h3>{{ $t('otpSettings.smsTitle') }}</h3>
              </div>
              <p class="otp-description">{{ $t('otpSettings.smsDescription') }}</p>
            </div>
          </template>
          <div class="countries-selector">
            <h4>{{ $t('otpSettings.countriesTitle') }}</h4>
            <p class="countries-description">{{ $t('otpSettings.countriesDescription') }}</p>
            <el-select
              v-model="smsSettings.selectedCountries"
              multiple
              filterable
              :placeholder="$t('otpSettings.selectCountries')"
              style="width: 100%"
              :loading="isSavingSmsCountries"
              @change="saveSmsCountries"
            >
              <el-option
                v-for="country in availableCountries"
                :key="country.code"
                :label="`${country.name} (${country.code})`"
                :value="country.code"
              />
            </el-select>
            <div v-if="smsSettings.selectedCountries.length > 0" class="selected-countries">
              <el-tag
                v-for="countryCode in smsSettings.selectedCountries"
                :key="countryCode"
                closable
                @close="removeSmsCountry(countryCode)"
                class="country-tag"
              >
                {{ getCountryName(countryCode) }}
              </el-tag>
            </div>
          </div>
        </ele-card>
        <ele-card class="otp-card">
          <template #header>
            <div class="card-header">
              <div class="otp-title">
                <el-icon class="otp-icon"><ChatDotRound /></el-icon>
                <h3>{{ $t('otpSettings.whatsappTitle') }}</h3>
              </div>
              <p class="otp-description">{{ $t('otpSettings.whatsappDescription') }}</p>
            </div>
          </template>
          <div class="countries-selector">
            <h4>{{ $t('otpSettings.countriesTitle') }}</h4>
            <p class="countries-description">{{ $t('otpSettings.countriesDescription') }}</p>
            <el-select
              v-model="whatsappSettings.selectedCountries"
              multiple
              filterable
              :placeholder="$t('otpSettings.selectCountries')"
              style="width: 100%"
              :loading="isSavingWhatsappCountries"
              @change="saveWhatsappCountries"
            >
              <el-option
                v-for="country in availableCountries"
                :key="country.code"
                :label="`${country.name} (${country.code})`"
                :value="country.code"
              />
            </el-select>
            <div v-if="whatsappSettings.selectedCountries.length > 0" class="selected-countries">
              <el-tag
                v-for="countryCode in whatsappSettings.selectedCountries"
                :key="countryCode"
                closable
                @close="removeWhatsappCountry(countryCode)"
                class="country-tag"
              >
                {{ getCountryName(countryCode) }}
              </el-tag>
            </div>
          </div>
        </ele-card>
      </div>
    </div>
  </ele-page>
</template>

<script>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import {
  Message,
  ChatDotRound
} from '@element-plus/icons-vue'

export default {
  name: 'OtpSettings',
  components: {
    Message,
    ChatDotRound
  },
  setup() {
    const isLoading = ref(false)
    const error = ref(null)
    const isSavingSmsCountries = ref(false)
    const isSavingWhatsappCountries = ref(false)

    const smsSettings = reactive({
      selectedCountries: ['US', 'GB', 'CA']
    })

    const whatsappSettings = reactive({
      selectedCountries: ['US', 'DE', 'FR', 'AU']
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

    const getCountryName = (code) => {
      const country = availableCountries.value.find(c => c.code === code)
      return country ? country.name : code
    }

    const saveSmsCountries = async () => {
      if (isSavingSmsCountries.value) return

      isSavingSmsCountries.value = true
      try {
        await mockApiCall(800)
        ElMessage.success('SMS countries updated successfully')
      } catch (err) {
        ElMessage.error('Failed to update SMS countries')
        console.error('Error saving SMS countries:', err)
      } finally {
        isSavingSmsCountries.value = false
      }
    }


    const saveWhatsappCountries = async () => {
      if (isSavingWhatsappCountries.value) return

      isSavingWhatsappCountries.value = true
      try {
        await mockApiCall(800)
        ElMessage.success('WhatsApp countries updated successfully')
      } catch (err) {
        ElMessage.error('Failed to update WhatsApp countries')
        console.error('Error saving WhatsApp countries:', err)
      } finally {
        isSavingWhatsappCountries.value = false
      }
    }

    const removeSmsCountry = (countryCode) => {
      const index = smsSettings.selectedCountries.indexOf(countryCode)
      if (index > -1) {
        smsSettings.selectedCountries.splice(index, 1)
        saveSmsCountries()
      }
    }
    const removeWhatsappCountry = (countryCode) => {
      const index = whatsappSettings.selectedCountries.indexOf(countryCode)
      if (index > -1) {
        whatsappSettings.selectedCountries.splice(index, 1)
        saveWhatsappCountries()
      }
    }
    const loadOtpSettings = async () => {
      isLoading.value = true
      error.value = null

      try {
        await mockApiCall(1500)
      } catch (err) {
        error.value = 'Failed to load OTP settings'
        console.error('Error loading OTP settings:', err)
      } finally {
        isLoading.value = false
      }
    }
    onMounted(() => {
      loadOtpSettings()
    })

    return {
      isLoading,
      error,
      isSavingSmsCountries,
      isSavingWhatsappCountries,
      smsSettings,
      whatsappSettings,
      availableCountries,
      getCountryName,
      saveSmsCountries,
      saveWhatsappCountries,
      removeSmsCountry,
      removeWhatsappCountry
    }
  }
}
</script>

<style scoped>
.otp-settings-container {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.loading-container,
.error-container {
  padding: 40px 20px;
}

.otp-content {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.otp-card {
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.otp-card:hover {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
}

.card-header {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.otp-title {
  display: flex;
  align-items: center;
  gap: 12px;
}

.otp-title h3 {
  margin: 0;
  font-size: 20px;
  font-weight: 600;
  color: #303133;
}

.otp-icon {
  font-size: 24px;
  color: #409eff;
}

.otp-description {
  margin: 0;
  color: #606266;
  font-size: 14px;
  line-height: 1.5;
}

.countries-selector {
  padding: 20px;
}

.countries-selector h4 {
  margin: 0 0 8px 0;
  font-size: 16px;
  font-weight: 500;
  color: #303133;
}

.countries-description {
  margin: 0 0 16px 0;
  color: #909399;
  font-size: 13px;
  line-height: 1.4;
}

.selected-countries {
  margin-top: 16px;
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.country-tag {
  margin: 0;
}

@media (max-width: 768px) {
  .otp-settings-container {
    padding: 16px;
  }

  .otp-title {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }

  .countries-selector {
    padding: 16px;
  }
}
</style>
