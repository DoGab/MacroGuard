<script lang="ts">
  import "../app.css";
  import DesktopNav from "$lib/components/navigation/DesktopNav.svelte";
  import MobileDock from "$lib/components/navigation/MobileDock.svelte";
  import AddEntryModal from "$lib/components/navigation/AddEntryModal.svelte";

  let { children } = $props();
  let addMenuOpen = $state(false);
</script>

<svelte:head>
  <link rel="preconnect" href="https://fonts.googleapis.com" />
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous" />
  <link
    href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap"
    rel="stylesheet"
  />
  <meta name="theme-color" content="#22c55e" />
  <meta name="apple-mobile-web-app-capable" content="yes" />
  <meta name="apple-mobile-web-app-status-bar-style" content="black-translucent" />
</svelte:head>

<div class="min-h-screen bg-base-100 flex flex-col">
  <!-- Desktop Header Navigation -->
  <DesktopNav bind:addMenuOpen />

  <!-- Add Entry Modal -->
  <AddEntryModal bind:open={addMenuOpen} />

  <!-- Main Content -->
  <main class="flex-1 pb-20 lg:pb-0">
    {@render children()}
  </main>

  <!-- Mobile Bottom Dock -->
  <MobileDock bind:addMenuOpen />
</div>
