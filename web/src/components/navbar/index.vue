<template>
  <div class="navbar">
    <a-modal   :visible="showRoleForm" title="切换角色" @ok="showRoleForm=false"  @cancel="showRoleForm=false">
      <a-space>
        <a-tag @click="handleSwitchRole(item.id)" :bordered="roleData.current.name===item.name" :style="{'cursor':'pointer'}" :color="roleData.current.name===item.name?'cyan':''" :key="idx" v-for="(item,idx) in roleData.selectable" >
          {{item.name}}
        </a-tag>
      </a-space>
    </a-modal>
    <div class="left-side">
      <a-space>
        <img alt="logo" class="logo-img" :src="logoImg" />
        <a-typography-title
          :style="{ margin: 0, fontSize: '18px' }"
          :heading="5"
        >
          Nexwind
        </a-typography-title>
        <icon-menu-fold
          v-if="!topMenu && appStore.device === 'mobile'"
          style="font-size: 22px; cursor: pointer"
          @click="toggleDrawerMenu"
        />
      </a-space>
    </div>
    <div class="center-side">
      <Menu v-if="topMenu" />
    </div>
    <ul class="right-side">
<!--      <li>-->
<!--        <a-tooltip :content="$t('settings.search')">-->
<!--          <a-button class="nav-btn" type="outline" :shape="'circle'">-->
<!--            <template #icon>-->
<!--              <icon-search />-->
<!--            </template>-->
<!--          </a-button>-->
<!--        </a-tooltip>-->
<!--      </li>-->
      <li>
<!--        <a-tooltip :content="$t('settings.language')">-->
<!--          <a-button-->
<!--            class="nav-btn"-->
<!--            type="outline"-->
<!--            :shape="'circle'"-->
<!--            @click="setDropDownVisible"-->
<!--          >-->
<!--            <template #icon>-->
<!--              <icon-language />-->
<!--            </template>-->
<!--          </a-button>-->
<!--        </a-tooltip>-->
        <a-dropdown trigger="click" @select="changeLocale as any">
          <div ref="triggerBtn" class="trigger-btn"></div>
          <template #content>
            <a-doption
              v-for="item in locales"
              :key="item.value"
              :value="item.value"
            >
              <template #icon>
                <icon-check v-show="item.value === currentLocale" />
              </template>
              {{ item.label }}
            </a-doption>
          </template>
        </a-dropdown>
      </li>
      <li>
        <a-tooltip
          :content="
            theme === 'light'
              ? $t('settings.navbar.theme.toDark')
              : $t('settings.navbar.theme.toLight')
          "
        >
          <a-button
            class="nav-btn"
            type="outline"
            :shape="'circle'"
            @click="handleToggleTheme"
          >
            <template #icon>
              <icon-moon-fill v-if="theme === 'dark'" />
              <icon-sun-fill v-else />
            </template>
          </a-button>
        </a-tooltip>
      </li>
<!--      <li>-->
<!--        <a-tooltip :content="$t('settings.navbar.alerts')">-->
<!--          <div class="message-box-trigger">-->
<!--            <a-badge :count="9" dot>-->
<!--              <a-button-->
<!--                class="nav-btn"-->
<!--                type="outline"-->
<!--                :shape="'circle'"-->
<!--                @click="setPopoverVisible"-->
<!--              >-->
<!--                <icon-notification />-->
<!--              </a-button>-->
<!--            </a-badge>-->
<!--          </div>-->
<!--        </a-tooltip>-->
<!--        <a-popover-->
<!--          trigger="click"-->
<!--          :arrow-style="{ display: 'none' }"-->
<!--          :content-style="{ padding: 0, minWidth: '400px' }"-->
<!--          content-class="message-popover"-->
<!--        >-->
<!--          <div ref="refBtn" class="ref-btn"></div>-->
<!--          <template #content>-->
<!--            <message-box />-->
<!--          </template>-->
<!--        </a-popover>-->
<!--      </li>-->
      <li>
        <a-tooltip
          :content="
            isFullscreen
              ? $t('settings.navbar.screen.toExit')
              : $t('settings.navbar.screen.toFull')
          "
        >
          <a-button
            class="nav-btn"
            type="outline"
            :shape="'circle'"
            @click="toggleFullScreen"
          >
            <template #icon>
              <icon-fullscreen-exit v-if="isFullscreen" />
              <icon-fullscreen v-else />
            </template>
          </a-button>
        </a-tooltip>
      </li>
      <li>
        <!--        <a-tooltip :content="$t('settings.title')">-->
        <!--          <a-button-->
        <!--            class="nav-btn"-->
        <!--            type="outline"-->
        <!--            :shape="'circle'"-->
        <!--            @click="setVisible"-->
        <!--          >-->
        <!--            <template #icon>-->
        <!--              <icon-settings />-->
        <!--            </template>-->
        <!--          </a-button>-->
        <!--        </a-tooltip>-->

        <a-dropdown trigger="click">
          <a-avatar
            :size="32"
            :style="{ marginRight: '8px', cursor: 'pointer' }"
          >
            <img alt="avatar" :src="avatar" />
          </a-avatar>
          <template #content>
            <a-doption>
              <a-space @click="switchRoles">
                <icon-tag />
                <span>
                  切换角色
                </span>
              </a-space>
            </a-doption>
            <a-doption>
              <a-space @click="$router.push({ name: 'Info' })">
                <icon-user />
                <span>
                  {{ $t('messageBox.userCenter') }}
                </span>
              </a-space>
            </a-doption>
            <a-doption>
              <a-space @click="$router.push({ name: 'Setting' })">
                <icon-settings />
                <span>
                  {{ $t('messageBox.userSettings') }}
                </span>
              </a-space>
            </a-doption>
            <a-doption>
              <a-space @click="handleLogout">
                <icon-export />
                <span>
                  {{ $t('messageBox.logout') }}
                </span>
              </a-space>
            </a-doption>
          </template>
        </a-dropdown>
      </li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
  import logoImg from '@/assets/images/logo_v2.png';
  import { computed, ref, inject } from 'vue';
  import { Message } from '@arco-design/web-vue';
  import { useDark, useToggle, useFullscreen } from '@vueuse/core';
  import { useAppStore, useUserStore } from '@/store';
  import { LOCALE_OPTIONS } from '@/locale';
  import useLocale from '@/hooks/locale';
  import useUser from '@/hooks/user';
  import {getRoles,switchRole} from "@/api/auth";
  import Menu from '@/components/menu/index.vue';
  import MessageBox from '../message-box/index.vue';

  interface RoleData {
    current:any,
    selectable:any
  }
  const roleData:RoleData={
    current:{},
    selectable:[]
  }
  const appStore = useAppStore();
  const userStore = useUserStore();
  const { logout } = useUser();
  const { changeLocale, currentLocale } = useLocale();
  const { isFullscreen, toggle: toggleFullScreen } = useFullscreen();
  const locales = [...LOCALE_OPTIONS];
  const avatar = computed(() => {
    return userStore.avatar;
  });
  const theme = computed(() => {
    return appStore.theme;
  });

  const topMenu = computed(() => appStore.topMenu && appStore.menu);
  const isDark = useDark({
    selector: 'body',
    attribute: 'arco-theme',
    valueDark: 'dark',
    valueLight: 'light',
    storageKey: 'arco-theme',
    onChanged(dark: boolean) {
      // overridden default behavior
      appStore.toggleTheme(dark);
    },
  });
  const toggleTheme = useToggle(isDark);
  const handleToggleTheme = () => {
    toggleTheme();
  };
  const setVisible = () => {
    appStore.updateSettings({ globalSettings: true });
  };
  const showRoleForm=ref(false);
  const handleSwitchRole=async (id:number)=>{
    const resp=await switchRole({
      prev: roleData.current.id,
      after:id
    });
    if(resp.code===20000){
      Message.success("切换成成功");
      showRoleForm.value=false;
      setTimeout(()=>{
        window.location.reload();
      },500)
    }else{
      Message.error(resp.message);
    }
  };
  const refBtn = ref();
  const triggerBtn = ref();
  const setPopoverVisible = () => {
    const event = new MouseEvent('click', {
      view: window,
      bubbles: true,
      cancelable: true,
    });
    refBtn.value.dispatchEvent(event);
  };
  const handleLogout = () => {
    logout();
  };
  const setDropDownVisible = () => {
    const event = new MouseEvent('click', {
      view: window,
      bubbles: true,
      cancelable: true,
    });
    triggerBtn.value.dispatchEvent(event);
  };


  const switchRoles = async () => {
    const resp=await getRoles()
    if(resp.code===20000){
      roleData.selectable=resp.data.selectable;
      roleData.current=resp.data.current;
    }
    showRoleForm.value=true;
  };
  const toggleDrawerMenu = inject('toggleDrawerMenu') as () => void;
</script>

<style scoped lang="less">
  .logo-img {
    width: 30px;
    height: 30px;
  }
  .navbar {
    display: flex;
    justify-content: space-between;
    height: 100%;
    background-color: var(--color-bg-2);
    border-bottom: 1px solid var(--color-border);
  }

  .left-side {
    display: flex;
    align-items: center;
    padding-left: 20px;
  }

  .center-side {
    flex: 1;
  }

  .right-side {
    display: flex;
    padding-right: 20px;
    list-style: none;
    :deep(.locale-select) {
      border-radius: 20px;
    }
    li {
      display: flex;
      align-items: center;
      padding: 0 10px;
    }

    a {
      color: var(--color-text-1);
      text-decoration: none;
    }
    .nav-btn {
      border-color: rgb(var(--gray-2));
      color: rgb(var(--gray-8));
      font-size: 16px;
    }
    .trigger-btn,
    .ref-btn {
      position: absolute;
      bottom: 14px;
    }
    .trigger-btn {
      margin-left: 14px;
    }
  }
</style>

<style lang="less">
  .message-popover {
    .arco-popover-content {
      margin-top: 0;
    }
  }
</style>
