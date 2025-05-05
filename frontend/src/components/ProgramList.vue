<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import {
  NAlert,
  NAvatar,
  NCard,
  NEmpty,
  NGrid,
  NGridItem,
  NPageHeader,
  NSpace,
  NSpin,
  NTag,
  NText,
  NThing
} from 'naive-ui'

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

function formatDate(d: string) {
  if (!d) return ''
  const date = new Date(d)
  const now = new Date()
  const diffTime = Math.abs(now.getTime() - date.getTime())
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))

  if (diffDays <= 1) return 'Today'
  if (diffDays <= 2) return 'Yesterday'
  if (diffDays <= 7) return `${diffDays} days ago`
  return date.toLocaleDateString(undefined, { month: 'short', day: 'numeric', year: 'numeric' })
}

// Generate a consistent color based on the program name
function getProgramColor(name: string) {
  const colors = [
    '#2080f0', '#18a058', '#f0a020', '#d03050',
    '#8a2be2', '#1890ff', '#52c41a', '#faad14',
    '#722ed1', '#eb2f96', '#f5222d', '#fa541c'
  ]
  let hash = 0
  for (let i = 0; i < name.length; i++) {
    hash = name.charCodeAt(i) + ((hash << 5) - hash)
  }
  hash = Math.abs(hash)
  return colors[hash % colors.length]
}

// Get initials from program name for avatar
function getInitials(name: string) {
  return name
      .split(' ')
      .map(word => word[0])
      .join('')
      .toUpperCase()
      .substring(0, 2)
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

      <n-grid
          v-else
          :x-gap="24"
          :y-gap="24"
          cols="auto-fit"
          min-cols-width="280px"
      >
        <n-grid-item v-for="program in programs" :key="program.id">
          <n-card
              hoverable
              class="program-card"
              @click="select(program.id)"
          >
            <n-thing>
              <template #avatar>
                <n-avatar
                    round
                    :style="{ backgroundColor: getProgramColor(program.name) }"
                >
                  {{ getInitials(program.name) }}
                </n-avatar>
              </template>

              <template #header>
                <n-text class="program-title">{{ program.name }}</n-text>
              </template>

              <template #description>
                <n-space vertical size="small">
                  <n-space align="center" size="small">
                    <n-icon size="16">
                      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24">
                        <path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8s8 3.59 8 8s-3.59 8-8 8zm.5-13H11v6l5.2 3.2l.8-1.3l-4.5-2.7V7z"/>
                      </svg>
                    </n-icon>
                    <n-text depth="3" class="program-date">
                      {{ formatDate(program.created_at || '') }}
                    </n-text>
                  </n-space>

                  <n-tag v-if="program.shared_by" size="small">
                    Shared by {{ program.shared_by }}
                  </n-tag>
                </n-space>
              </template>
            </n-thing>
          </n-card>
        </n-grid-item>
      </n-grid>
    </n-space>
  </div>
</template>

<style scoped>
.program-list {
  width: 100%;
}

.content-space {
  margin-top: 2rem;
  width: 100%;
}

.program-card {
  position: relative;
  overflow: hidden;
  border-radius: 12px;
  transition: all 0.3s ease;
  height: 100%;
}

.program-card:hover {
  transform: translateY(-6px);
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.1);
}

.program-title {
  font-size: 1.1rem;
  font-weight: 600;
  line-height: 1.5;
}

.program-date {
  font-size: 0.85rem;
}
</style>