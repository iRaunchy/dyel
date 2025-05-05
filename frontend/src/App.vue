<script setup lang="ts">
import {RouterLink, RouterView, useRoute} from 'vue-router'
import {computed, ref} from 'vue'
import { darkTheme } from 'naive-ui'


const route = useRoute()
const currentPath = computed(() => route.path)
const isDarkMode = ref(false)

function toggleDarkMode() {
  isDarkMode.value = !isDarkMode.value
}
</script>

<template>
  <n-config-provider :theme="isDarkMode ? darkTheme : null">
    <n-message-provider>
      <n-layout position="absolute">
        <!-- Header with navigation -->
        <n-layout-header bordered class="header">
          <div class="container header-content">
            <div class="logo">
              DoYouEven<strong>LIFT</strong>
            </div>

            <n-space>
              <n-button
                  text
                  :type="currentPath === '/programs' ? 'primary' : 'default'"
                  class="nav-button"
              >
                <template #icon>
                  <n-icon>
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24">
                      <path fill="currentColor" d="M4 13h6V5H4zm8 0h8V5h-8zM4 21h6v-6H4zm8 0h8v-6h-8z"/>
                    </svg>
                  </n-icon>
                </template>
                <RouterLink to="/programs">Programs</RouterLink>
              </n-button>

              <n-button
                  text
                  :type="currentPath === '/programs/new' ? 'primary' : 'default'"
                  class="nav-button"
              >
                <template #icon>
                  <n-icon>
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24">
                      <path fill="currentColor" d="M18 13h-5v5h-2v-5H6v-2h5V6h2v5h5z"/>
                    </svg>
                  </n-icon>
                </template>
                <RouterLink to="/programs/new">Create New</RouterLink>
              </n-button>

              <!-- Theme toggle button -->
              <n-button circle @click="toggleDarkMode">
                <template #icon>
                  <n-icon>
                    <svg v-if="isDarkMode" xmlns="http://www.w3.org/2000/svg" width="16" height="16"
                         viewBox="0 0 24 24">
                      <path fill="currentColor"
                            d="M12 9c1.65 0 3 1.35 3 3s-1.35 3-3 3s-3-1.35-3-3s1.35-3 3-3m0-2c-2.76 0-5 2.24-5 5s2.24 5 5 5s5-2.24 5-5s-2.24-5-5-5zM2 13h2c.55 0 1-.45 1-1s-.45-1-1-1H2c-.55 0-1 .45-1 1s.45 1 1 1zm18 0h2c.55 0 1-.45 1-1s-.45-1-1-1h-2c-.55 0-1 .45-1 1s.45 1 1 1zM11 2v2c0 .55.45 1 1 1s1-.45 1-1V2c0-.55-.45-1-1-1s-1 .45-1 1zm0 18v2c0 .55.45 1 1 1s1-.45 1-1v-2c0-.55-.45-1-1-1s-1 .45-1 1zM5.99 4.58c-.39-.39-1.03-.39-1.41 0c-.39.39-.39 1.03 0 1.41l1.06 1.06c.39.39 1.03.39 1.41 0s.39-1.03 0-1.41L5.99 4.58zm12.37 12.37c-.39-.39-1.03-.39-1.41 0c-.39.39-.39 1.03 0 1.41l1.06 1.06c.39.39 1.03.39 1.41 0c.39-.39.39-1.03 0-1.41l-1.06-1.06zm1.06-10.96c.39-.39.39-1.03 0-1.41c-.39-.39-1.03-.39-1.41 0l-1.06 1.06c-.39.39-.39 1.03 0 1.41s1.03.39 1.41 0l1.06-1.06zM7.05 18.36c.39-.39.39-1.03 0-1.41c-.39-.39-1.03-.39-1.41 0l-1.06 1.06c-.39.39-.39 1.03 0 1.41s1.03.39 1.41 0l1.06-1.06z"/>
                    </svg>
                    <svg v-else xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24">
                      <path fill="currentColor"
                            d="M12 3c-4.97 0-9 4.03-9 9s4.03 9 9 9s9-4.03 9-9c0-.46-.04-.92-.1-1.36c-.98 1.37-2.58 2.26-4.4 2.26c-2.98 0-5.4-2.42-5.4-5.4c0-1.81.89-3.42 2.26-4.4c-.44-.06-.9-.1-1.36-.1z"/>
                    </svg>
                  </n-icon>
                </template>
              </n-button>
            </n-space>
          </div>
        </n-layout-header>

        <!-- Main Content -->
        <n-layout-content class="content">
          <div class="container">
            <n-card class="content-card">
              <RouterView/>
            </n-card>
          </div>
        </n-layout-content>

        <!-- Footer -->
        <n-layout-footer bordered class="footer">
          <div class="container">
            <n-space justify="center">
              <n-text depth="3">
                DoYouEvenLift © {{ new Date().getFullYear() }}
              </n-text>
              <n-divider vertical/>
              <n-text depth="3">
                <a href="#" class="footer-link">Privacy</a>
              </n-text>
              <n-divider vertical/>
              <n-text depth="3">
                <a href="#" class="footer-link">Terms</a>
              </n-text>
            </n-space>
          </div>
        </n-layout-footer>
      </n-layout>
    </n-message-provider>
  </n-config-provider>
</template>

<style scoped>
/* Container for consistent width */
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 16px;
  width: 100%;
}

/* Header styles */
.header {
  backdrop-filter: blur(10px);
  background-color: rgba(255, 255, 255, 0.7);
  position: sticky;
  top: 0;
  z-index: 100;
  height: 64px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100%;
}

.logo {
  font-size: 1.25rem;
  letter-spacing: -0.5px;
}

.nav-button {
  padding: 8px 16px;
  font-weight: 500;
}

/* Navigation link styles */
:deep(a) {
  text-decoration: none;
  color: inherit;
}

/* Content area styles */
.content {
  padding: 32px 0;
  min-height: calc(100vh - 64px - 60px); /* viewport height minus header and footer */
  background-color: var(--n-color-modal-backdrop);
}

.content-card {
  min-height: 300px;
}

/* Footer styles */
.footer {
  padding: 16px 0;
  height: 60px;
  text-align: center;
}

.footer-link {
  color: inherit;
  text-decoration: none;
  transition: opacity 0.2s;
}

.footer-link:hover {
  opacity: 0.8;
}

/* Dark mode tweaks */
:deep(.n-config-provider.n-config-provider--dark) {
  .header {
    background-color: rgba(38, 40, 48, 0.7);
  }

  .footer {
    background-color: rgba(38, 40, 48, 0.8);
    border-top-color: rgba(255, 255, 255, 0.08);
  }

  .footer :deep(.n-text) {
    color: rgba(255, 255, 255, 0.7);
  }

  .footer :deep(.n-divider) {
    background-color: rgba(255, 255, 255, 0.15);
  }
}
</style>