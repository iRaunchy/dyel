<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import {
  NAlert,
  NAvatar,
  NEmpty,
  NPageHeader,
  NSpace,
  NSpin,
} from 'naive-ui'
import { formatDate, getProgramColor, getInitials } from '../utils/utils'

interface Program {
  id: string
  name: string
  shared_by?: string
  created_at?: string
}

const programs = ref<Program[]>([])
const loading = ref(false)
const error = ref('')
const router = useRouter()

function select(id: string) {
  router.push(`/programs/${id}`)
}

async function loadPrograms() {
  loading.value = true
  error.value = ''
  try {
    const res = await fetch('/api/v1/programs')
    if (!res.ok) throw new Error(res.statusText)
    programs.value = await res.json()
  } catch (err: any) {
    error.value = err.message || 'Failed to load programs'
  } finally {
    loading.value = false
  }
}

onMounted(loadPrograms)
</script>

<template>
  <div class="program-list">
    <n-page-header
        title="My Workout Programs"
        subtitle="Browse and manage your workout routines"
    />

    <n-space vertical align="center" size="large" class="content-space">
      <n-spin v-if="loading" size="large"/>

      <n-alert
          v-else-if="error"
          type="error"
          show-icon
          :title="error"
      />

      <n-empty
          v-else-if="!programs.length"
          description="No programs yet. Create one to get started!"
      >
        <template #extra>
          <n-button type="primary" @click="$router.push('/programs/new')">
            Create Your First Program
          </n-button>
        </template>
      </n-empty>

      <div v-else class="programs-grid">
        <div
            v-for="program in programs"
            :key="program.id"
            class="glass-card"
            @click="select(program.id)"
        >
          <div class="card-content">
            <div class="card-top">
              <n-avatar
                  round
                  :style="{ backgroundColor: getProgramColor(program.name) }"
                  class="avatar-glass"
              >
                {{ getInitials(program.name) }}
              </n-avatar>

              <div class="program-title">{{ program.name }}</div>
            </div>

            <div class="card-bottom">
              <div class="date-info">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" class="time-icon">
                  <path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8s8 3.59 8 8s-3.59 8-8 8zm.5-13H11v6l5.2 3.2l.8-1.3l-4.5-2.7V7z"/>
                </svg>
                <span class="program-date">
                  {{ formatDate(program.created_at || '') }}
                </span>
              </div>

              <div v-if="program.shared_by" class="shared-info">
                Shared by {{ program.shared_by }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </n-space>
  </div>
</template>

<style>
/* Global styles to ensure text is visible */
[class*="n-tag__content"] {
  color: rgba(0, 0, 0, 0.65) !important;
}
</style>

<style scoped>
.program-list {
  width: 100%;
  max-width: 1400px;
  margin: 0 auto;
}

.content-space {
  margin-top: 2rem;
  width: 100%;
}

.programs-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(275px, 1fr));
  gap: 24px;
  width: 100%;
}

/* For larger screens, limit to 4 cards per row max */
@media (min-width: 1400px) {
  .programs-grid {
    grid-template-columns: repeat(4, 1fr);
  }
}

/* Adjustments for smaller screens */
@media (max-width: 600px) {
  .programs-grid {
    grid-template-columns: 1fr;
  }
}

.glass-card {
  position: relative;
  overflow: hidden;
  border-radius: 16px;
  background: rgba(255, 255, 255, 0.7) !important;
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  border: 1px solid rgba(0, 0, 0, 0.1);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.07);
  transition: all 0.3s ease;
  cursor: pointer;
  height: 200px; /* Increased height to fit content */
}

.glass-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 50%;
  background: linear-gradient(
      to bottom,
      rgba(255, 255, 255, 0.9),
      rgba(255, 255, 255, 0.5)
  );
  border-radius: 16px 16px 0 0;
  z-index: 0;
}

.card-content {
  position: relative;
  z-index: 1;
  padding: 18px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  height: 100%;
  box-sizing: border-box;
}

.card-top {
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 110px; /* Increased height for the top section */
  overflow: hidden;
}

.avatar-glass {
  margin-bottom: 10px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  border: 2px solid rgba(0, 0, 0, 0.1);
  flex-shrink: 0;
}

.program-title {
  font-size: 1.1rem;
  font-weight: 600;
  line-height: 1.3;
  color: #000000 !important;
  text-align: center;
  word-break: break-word;
  max-height: 2.6rem; /* Limit to 2 lines */
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  margin: 0; /* Remove default margins */
}

.card-bottom {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  margin-top: auto;
  min-height: 50px; /* Ensure bottom section has enough room */
}

.date-info {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
}

.time-icon {
  color: rgba(0, 0, 0, 0.6);
}

.program-date {
  font-size: 0.85rem;
  color: rgba(0, 0, 0, 0.6) !important;
}

.shared-info {
  background: rgba(255, 255, 255, 0.6);
  border: 1px solid rgba(0, 0, 0, 0.05);
  border-radius: 4px;
  padding: 3px 8px;
  font-size: 0.8rem;
  color: rgba(0, 0, 0, 0.55) !important;
  text-align: center;
  max-width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.glass-card:hover {
  transform: translateY(-6px);
  box-shadow: 0 14px 40px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(0, 0, 0, 0.15);
}
</style>