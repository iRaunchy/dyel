<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  NAvatar,
  NButton,
  NCard,
  NDivider,
  NEmpty,
  NGrid,
  NGridItem,
  NIcon,
  NPageHeader,
  NSpace,
  NSpin,
  NStatistic,
  NTimeline,
  NTimelineItem,
  NAlert,
  NTag,
  NText
} from 'naive-ui'

interface Exercise {
  id?: string
  name: string
}

interface Day {
  id: string
  name: string
  exercises: Exercise[] | null
}

interface Program {
  id: string
  name: string
  shared_by: string
  days: Day[]
  created_at: string
  updated_at: string
}

const route = useRoute()
const router = useRouter()
const program = ref<Program | null>(null)
const loading = ref(false)
const error = ref('')

// Get day emoji based on name
function getDayEmoji(dayName: string): string {
  const emojis: Record<string, string> = {
    'Monday': '🏋️',
    'Tuesday': '💪',
    'Wednesday': '🔄',
    'Thursday': '⚡',
    'Friday': '🏆',
    'Saturday': '🔥',
    'Sunday': '🧘'
  }
  return emojis[dayName] || '📅'
}

// Calculate program stats
const stats = computed(() => {
  if (!program.value) return { days: 0, exercises: 0 }

  const days = program.value.days.length
  let exerciseCount = 0

  program.value.days.forEach(day => {
    exerciseCount += day.exercises?.length || 0
  })

  return { days, exercises: exerciseCount }
})

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

// Format date nicely
function formatDate(dateString: string): string {
  const date = new Date(dateString)
  return date.toLocaleDateString(undefined, {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

async function loadProgram() {
  const id = route.params.id
  if (!id) {
    error.value = 'No program ID provided'
    return
  }
  loading.value = true
  error.value = ''
  try {
    const res = await fetch(`/api/v1/programs/${id}`)
    if (!res.ok) throw new Error(res.statusText)
    program.value = await res.json()
  } catch (err: any) {
    error.value = err.message || 'Failed to load program'
  } finally {
    loading.value = false
  }
}

function handleEdit() {
  // You can implement program editing in the future
  alert('Edit functionality not implemented yet')
}

function handleDelete() {
  if (!confirm('Are you sure you want to delete this program?')) return

  const id = route.params.id
  fetch(`/api/v1/programs/${id}`, {
    method: 'DELETE'
  }).then(res => {
    if (res.ok) {
      router.push('/programs')
    } else {
      alert('Failed to delete program')
    }
  })
}

onMounted(loadProgram)
</script>

<template>
  <div class="program-detail">
    <n-spin :show="loading">
      <n-space vertical size="large">
        <!-- Back button and header -->
        <n-page-header
            @back="router.push('/programs')"
            :title="program?.name || 'Program Details'"
            :subtitle="program ? `Created by ${program.shared_by}` : ''"
        >
          <template #avatar v-if="program">
            <n-avatar
                round
                :style="{ backgroundColor: getProgramColor(program.name) }"
                size="large"
            >
              {{ getInitials(program.name) }}
            </n-avatar>
          </template>
          <template #extra v-if="program">
            <n-space>
              <n-button @click="handleEdit">Edit</n-button>
              <n-button type="error" @click="handleDelete">Delete</n-button>
            </n-space>
          </template>
        </n-page-header>

        <!-- Error message -->
        <n-alert v-if="error" type="error" :title="error" />

        <!-- Program content -->
        <template v-if="program && !loading && !error">
          <!-- Program stats -->
          <n-card title="Program Overview" size="large">
            <n-grid :cols="2" :x-gap="16">
              <n-grid-item>
                <n-statistic label="Workout Days" :value="stats.days">
                  <template #prefix>
                    <n-icon>
                      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24">
                        <path fill="currentColor" d="M19 19H5V8h14m-3-7v2H8V1H6v2H5c-1.11 0-2 .89-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-1V1"/>
                      </svg>
                    </n-icon>
                  </template>
                </n-statistic>
              </n-grid-item>
              <n-grid-item>
                <n-statistic label="Total Exercises" :value="stats.exercises">
                  <template #prefix>
                    <n-icon>
                      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24">
                        <path fill="currentColor" d="M20.57 14.86L22 13.43L20.57 12L17 15.57L8.43 7L12 3.43L10.57 2L9.14 3.43L7.71 2L5.57 4.14L4.14 2.71L2.71 4.14l1.43 1.43L2 7.71l1.43 1.43L2 10.57L3.43 12L7 8.43L15.57 17L12 20.57L13.43 22l1.43-1.43L16.29 22l2.14-2.14l1.43 1.43l1.43-1.43l-1.43-1.43L22 16.29z"/>
                      </svg>
                    </n-icon>
                  </template>
                </n-statistic>
              </n-grid-item>
            </n-grid>

            <n-space vertical size="small" style="margin-top: 16px">
              <n-text depth="3">
                Created: {{ formatDate(program.created_at) }}
              </n-text>
              <n-text depth="3" v-if="program.created_at !== program.updated_at">
                Last Updated: {{ formatDate(program.updated_at) }}
              </n-text>
            </n-space>
          </n-card>

          <!-- Workout schedule -->
          <n-card title="Weekly Schedule" size="large">
            <template #header-extra>
              <n-icon size="18">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
                  <path fill="currentColor" d="M9 11H7v2h2v-2zm4 0h-2v2h2v-2zm4 0h-2v2h2v-2zm2-7h-1V2h-2v2H8V2H6v2H5c-1.11 0-1.99.9-1.99 2L3 20a2 2 0 0 0 2 2h14c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm0 16H5V9h14v11z"/>
                </svg>
              </n-icon>
            </template>

            <n-empty
                v-if="!program.days.length"
                description="No workout days scheduled"
            >
              <template #icon>
                <n-icon size="48">
                  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
                    <path fill="currentColor" d="M19 19H5V8h14m-3-7v2H8V1H6v2H5c-1.11 0-2 .89-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-1V1m-1 11h-5v5h5v-5Z"/>
                  </svg>
                </n-icon>
              </template>
            </n-empty>

            <n-grid :cols="{ xs: 1, sm: 2, md: 3, lg: 4 }" :x-gap="16" :y-gap="16" v-else>
              <n-grid-item v-for="day in program.days" :key="day.id">
                <n-card :title="day.name" size="small" class="day-card">
                  <template #header-extra>
                    <div class="day-emoji">{{ getDayEmoji(day.name) }}</div>
                  </template>

                  <n-empty
                      v-if="!day.exercises?.length"
                      description="No exercises scheduled"
                      size="small"
                  />

                  <n-timeline v-else>
                    <n-timeline-item
                        v-for="(exercise, idx) in day.exercises"
                        :key="idx"
                        :title="`${idx + 1}. ${typeof exercise === 'string' ? exercise : exercise.name}`"
                        :content="''"
                        :time="''"
                    >
                      <template #icon>
                        <n-icon>
                          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24">
                            <path fill="currentColor" d="M20.57 14.86L22 13.43L20.57 12L17 15.57L8.43 7L12 3.43L10.57 2L9.14 3.43L7.71 2L5.57 4.14L4.14 2.71L2.71 4.14l1.43 1.43L2 7.71l1.43 1.43L2 10.57L3.43 12L7 8.43L15.57 17L12 20.57L13.43 22l1.43-1.43L16.29 22l2.14-2.14l1.43 1.43l1.43-1.43l-1.43-1.43L22 16.29z"/>
                          </svg>
                        </n-icon>
                      </template>
                    </n-timeline-item>
                  </n-timeline>
                </n-card>
              </n-grid-item>
            </n-grid>
          </n-card>
        </template>
      </n-space>
    </n-spin>
  </div>
</template>

<style scoped>
.program-detail {
  width: 100%;
}

.day-card {
  height: 100%;
  transition: all 0.3s ease;
}

.day-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.09);
}

.day-emoji {
  font-size: 1.25rem;
}

:deep(.n-statistic-value) {
  font-size: 2rem !important;
  font-weight: 600;
}

:deep(.n-timeline-item-content) {
  min-height: 0;
}

:deep(.n-timeline-item:last-child .n-timeline-item-timeline) {
  height: 0;
}

:deep(.n-timeline) {
  padding-left: 0;
}

:deep(.n-timeline-item-content__content) {
  min-height: 0;
}

:deep(.n-timeline-item-content) {
  margin-top: -8px;
  margin-bottom: 10px;
}

:deep(.n-timeline-item-timeline) {
  top: 6px;
}

:deep(.n-timeline-item-content__title) {
  font-weight: 500;
}
</style>