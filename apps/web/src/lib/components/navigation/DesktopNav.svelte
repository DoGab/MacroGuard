<script lang="ts">
  import { Home, Clock, MessageCircle, User, Plus } from "lucide-svelte";
  import { browser } from "$app/environment";
  import ThemeSwap from "$lib/components/ui/ThemeSwap.svelte";
  import FontThemeSwap from "$lib/components/ui/FontThemeSwap.svelte";
  import logoDark from "$lib/assets/logo/logo_text_dark.svg";
  import logoLight from "$lib/assets/logo/logo_text_light.svg";

  let { addMenuOpen = $bindable() } = $props();

  // Track if current theme is dark
  let isDarkTheme = $state(true);

  // Determine which logo to use based on theme
  let currentLogo = $derived(isDarkTheme ? logoLight : logoDark);

  // Watch for theme changes on the html element
  $effect(() => {
    if (browser) {
      const html = document.documentElement;
      const updateTheme = () => {
        const theme = html.getAttribute("data-theme");
        isDarkTheme = theme === "darkorganic" || theme === "biohacker";
      };

      // Initial check
      updateTheme();

      // Watch for changes
      const observer = new MutationObserver(updateTheme);
      observer.observe(html, { attributes: true, attributeFilter: ["data-theme"] });

      return () => observer.disconnect();
    }
  });

  const navLinks = [
    { href: "/", label: "Home", icon: Home },
    { href: "/history", label: "History", icon: Clock },
    { href: "/chat", label: "Chat", icon: MessageCircle },
    { href: "/profile", label: "Profile", icon: User }
  ];
</script>

<header
  class="navbar bg-base-100/90 backdrop-blur-md sticky top-0 z-50 border-b border-base-200 shadow-sm"
>
  <div class="navbar-start">
    <a href="/" class="btn btn-ghost text-xl font-bold gap-2">
      <img src={currentLogo} alt="VitalStack Logo" class="w-45 h-45" />
    </a>
  </div>

  <!-- Desktop Nav Links -->
  <div class="navbar-center hidden lg:flex">
    <ul class="menu menu-horizontal px-1 gap-1">
      {#each navLinks as link (link.href)}
        {@const Icon = link.icon}
        <li>
          <a href={link.href} class="btn btn-ghost btn-sm gap-2">
            <Icon class="w-4 h-4" />
            {link.label}
          </a>
        </li>
      {/each}
    </ul>
  </div>

  <div class="navbar-end gap-2">
    <!-- Font Theme Swap -->
    <FontThemeSwap />
    <!-- Color Theme Swap -->
    <ThemeSwap />

    <!-- Desktop: Add button -->
    <div class="hidden lg:flex">
      <button class="btn btn-primary btn-sm gap-2" onclick={() => (addMenuOpen = !addMenuOpen)}>
        <Plus class="w-4 h-4" />
        Add
      </button>
    </div>
  </div>
</header>
