<script setup lang="ts">
import { reactive, toRefs, watch, ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import {
  NAlert,
  NButton,
  NCard,
  NCheckbox,
  NCheckboxGroup,
  NDivider,
  NForm,
  NFormItem,
  NGrid,
  NGridItem,
  NIcon,
  NInput,
  NInputGroup,
  NPageHeader,
  NSpace,
  NText,
  NCollapse,
  NCollapseItem,
  NEmpty,
  NSpin
} from 'naive-ui'

const dayOptions = [
  'Monday', 'Tuesday', 'Wednesday',
  'Thursday', 'Friday', 'Saturday', 'Sunday'
]

interface CreateForm {
  name: string
  sharedBy: string
  days: string[]
}

interface Errors {
  name: string
  sharedBy: string
  days: string
}

const router = useRouter()
const formRef = ref(null)

const state = reactive({
  form: {
    name: '',
    sharedBy: '',
    days: []
  } as CreateForm,

  // holds exercises per day name
  dayExercises: {} as Record<string, string[]>,

  errors: {
    name: '',
    sharedBy: '',
    days: ''
  } as Errors,

  apiError: '',
  submitting: false,
  isSaving: false
})

const hasExercises = computed(() => {
  return Object.values(state.dayExercises).some(exercises => exercises.length > 0)
})

watch(
    () => state.form.days.slice(),
    (
        newDays: string[],
        oldDays: string[] = []
    ) => {
      newDays.forEach(day => {
        if (!state.dayExercises[day]) {
          state.dayExercises[day] = []
        }
      })
      oldDays.forEach(day => {
        if (!newDays.includes(day)) {
          delete state.dayExercises[day]
        }
      })
    },
    { immediate: true }
)

function validate() {
  state.errors.name = state.form.name.trim() ? '' : 'Name is required'
  state.errors.sharedBy = state.form.sharedBy.trim() ? '' : 'Author name is required'
  state.errors.days = state.form.days.length ? '' : 'Select at least one day'
  return !state.errors.name && !state.errors.sharedBy && !state.errors.days
}

function addExercise(day: string) {
  state.dayExercises[day].push('')
}

function removeExercise(day: string, idx: number) {
  state.dayExercises[day].splice(idx, 1)
}

async function onSubmit() {
  if (!validate()) return

  state.submitting = true
  state.isSaving = true
  state.apiError = ''

  const payload = {
    name: state.form.name,
    shared_by: state.form.sharedBy,
    days: state.form.days.map(day => ({
      name: day,
      exercises: state.dayExercises[day].filter(ex => ex.trim()).map(ex => ({
        name: ex
      }))
    }))
  }

  try {
    const res = await fetch('/api/v1/programs', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(payload)
    })
    if (!res.ok) {
      const text = await res.text()
      throw new Error(text || res.statusText)
    }
    const created = await res.json() as { id?: string }
    if (!created.id) {
      throw new Error('Server did not return a new program ID')
    }
    await router.push(`/programs/${created.id}`)
  } catch (err: any) {
    state.apiError = err.message || 'Failed to create program'
    state.isSaving = false
  } finally {
    state.submitting = false
  }
}

const { form, errors, apiError, submitting, dayExercises, isSaving } = toRefs(state)
</script>

<template>
  <div class="program-create">
    <n-space vertical size="large">
      <n-page-header
          title="Create New Program"
          subtitle="Design your workout routine"
      >
        <template #extra>
          <n-space>
            <n-button @click="router.push('/programs')">
              Cancel
            </n-button>
            <n-button
                type="primary"
                :loading="submitting"
                :disabled="!form.name || !form.sharedBy || !form.days.length"
                @click="onSubmit"
            >
              Save Program
            </n-button>
          </n-space>
        </template>
      </n-page-header>

      <n-alert v-if="apiError" type="error" :title="apiError" />

      <n-spin :show="isSaving">
        <n-grid :cols="24" :x-gap="24">
          <!-- Left column - Form inputs -->
          <n-grid-item :span="24" :lg-span="8">
            <n-card title="Program Details" size="large">
              <n-form ref="formRef" :model="form" label-placement="left" label-width="100px" require-mark-placement="right-hanging">
                <n-form-item label="Name" :validation-status="errors.name ? 'error' : undefined" :feedback="errors.name">
                  <n-input v-model:value="form.name" placeholder="My Workout Program" />
                </n-form-item>

                <n-form-item label="Author" :validation-status="errors.sharedBy ? 'error' : undefined" :feedback="errors.sharedBy">
                  <n-input v-model:value="form.sharedBy" placeholder="Your name" />
                </n-form-item>

                <n-form-item label="Schedule" :validation-status="errors.days ? 'error' : undefined" :feedback="errors.days">
                  <n-checkbox-group v-model:value="form.days">
                    <n-space vertical>
                      <n-checkbox v-for="day in dayOptions" :key="day" :value="day" :label="day" />
                    </n-space>
                  </n-checkbox-group>
                </n-form-item>
              </n-form>
            </n-card>
          </n-grid-item>

          <!-- Right column - Exercises -->
          <n-grid-item :span="24" :lg-span="16">
            <n-card title="Exercises by Day" size="large">
              <template #header-extra>
                <n-text depth="3">{{ Object.keys(dayExercises).length }} days selected</n-text>
              </template>

              <n-empty
                  v-if="!form.days.length"
                  description="Select days from the left panel to add exercises"
              >
                <template #icon>
                  <n-icon>
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
                      <path fill="currentColor" d="M19 19H5V8h14m-3-7v2H8V1H6v2H5c-1.11 0-2 .89-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-1V1m-1 11h-5v5h5v-5Z"/>
                    </svg>
                  </n-icon>
                </template>
              </n-empty>

              <div v-else>
                <n-collapse arrow-placement="right">
                  <n-collapse-item
                      v-for="day in form.days"
                      :key="day"
                      :title="day"
                      :name="day"
                  >
                    <template #header-extra>
                      <n-text depth="3">{{ dayExercises[day]?.length || 0 }} exercises</n-text>
                    </template>

                    <n-space vertical>
                      <div v-if="dayExercises[day]?.length" class="exercise-list">
                        <div v-for="(exercise, idx) in dayExercises[day]" :key="`${day}-${idx}`" class="exercise-item">
                          <n-input-group>
                            <span class="exercise-number">{{ idx + 1 }}</span>
                            <n-input
                                v-model:value="dayExercises[day][idx]"
                                placeholder="Exercise name (e.g., Bench Press 3×10)"
                            />
                            <n-button quaternary circle @click="removeExercise(day, idx)">
                              <template #icon>
                                <n-icon>
                                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24">
                                    <path fill="currentColor" d="M19 6.41L17.59 5L12 10.59L6.41 5L5 6.41L10.59 12L5 17.59L6.41 19L12 13.41L17.59 19L19 17.59L13.41 12L19 6.41z"/>
                                  </svg>
                                </n-icon>
                              </template>
                            </n-button>
                          </n-input-group>
                        </div>
                      </div>

                      <n-space justify="center">
                        <n-button
                            @click="addExercise(day)"
                            secondary
                            strong
                            dashed
                            size="small"
                        >
                          <template #icon>
                            <n-icon>
                              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24">
                                <path fill="currentColor" d="M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z"/>
                              </svg>
                            </n-icon>
                          </template>
                          Add Exercise
                        </n-button>
                      </n-space>
                    </n-space>
                  </n-collapse-item>
                </n-collapse>

                <n-space justify="center" style="margin-top: 24px">
                  <n-text v-if="!hasExercises" depth="3">
                    No exercises added yet. Click "Add Exercise" for each day to build your program.
                  </n-text>
                  <n-button v-else type="primary" :loading="submitting" @click="onSubmit">Save Program</n-button>
                </n-space>
              </div>
            </n-card>
          </n-grid-item>
        </n-grid>
      </n-spin>
    </n-space>
  </div>
</template>

<style scoped>
.program-create {
  width: 100%;
}

.exercise-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.exercise-item {
  display: flex;
  align-items: center;
}

.exercise-number {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  padding: 6px 0;
  background-color: var(--n-color-info-suppl);
  color: var(--n-text-color-base);
  font-weight: 500;
  border-radius: 4px 0 0 4px;
  margin-right: -1px;
  border: 1px solid var(--n-border-color);
  border-right: none;
}

:deep(.n-input-group .n-input) {
  flex: 1;
}

:deep(.n-input-group .n-input .n-input__placeholder) {
  color: var(--n-text-color-3);
}

:deep(.n-checkbox-group .n-space:not(.n-space--inline)) {
  margin-top: 8px;
}
</style>