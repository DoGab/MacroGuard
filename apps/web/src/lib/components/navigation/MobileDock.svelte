<script lang="ts">
  import { Home, Clock, MessageCircle, User, Plus } from "lucide-svelte";

  let { addMenuOpen = $bindable() } = $props();
</script>

<div class="dock dock-md lg:hidden bg-base-100/90 backdrop-blur-md border-t border-base-200">
  <a href="/">
    <Home class="size-[1.2em]" />
    <span class="dock-label">Home</span>
  </a>

  <a href="/history">
    <Clock class="size-[1.2em]" />
    <span class="dock-label">History</span>
  </a>

  <!-- Center Add Button (elevated) -->
  <button
    class="!bg-primary !text-primary-content !rounded-full !w-14 !h-14 -translate-y-3 shadow-lg active:scale-95 transition-transform"
    onclick={() => (addMenuOpen = !addMenuOpen)}
    aria-label="Add entry"
  >
    <Plus class="w-7 h-7" />
  </button>

  <a href="/chat">
    <MessageCircle class="size-[1.2em]" />
    <span class="dock-label">Chat</span>
  </a>

  <a href="/profile">
    <User class="size-[1.2em]" />
    <span class="dock-label">Profile</span>
  </a>
</div>
