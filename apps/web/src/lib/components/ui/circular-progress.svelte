<script lang="ts">
  import type { Snippet } from "svelte";
  import { cn } from "$lib/utils";

  interface Props {
    size?: number; // ViewBox size, e.g., 64 or 120
    radius?: number; // Circle radius, e.g., 28 or 54
    strokeWidth?: number; // e.g., 5 or 8
    percent?: number; // Current progress percentage (0-100)
    addedPercent?: number; // Additional progress percentage (0-100)
    class?: string; // Container class
    children?: Snippet; // Inner content (like flame icon or text)
  }

  let {
    size = 64,
    radius = 28,
    strokeWidth = 5,
    percent = 0,
    addedPercent = 0,
    class: className,
    children
  }: Props = $props();

  let center = $derived(size / 2);
  let circumference = $derived(2 * Math.PI * radius);

  let currentPercent = $derived(Math.min(Math.max(percent, 0), 100));
  let added = $derived(Math.min(Math.max(addedPercent, 0), 100 - currentPercent));

  let offsetCurrent = $derived(circumference - (currentPercent / 100) * circumference);
  let offsetCombined = $derived(circumference - ((currentPercent + added) / 100) * circumference);
</script>

<div class={cn("relative shrink-0 flex items-center justify-center", className)}>
  <svg class="w-full h-full -rotate-90" viewBox="0 0 {size} {size}">
    <!-- Background -->
    <circle
      cx={center}
      cy={center}
      r={radius}
      fill="none"
      stroke="currentColor"
      stroke-width={strokeWidth}
      class="text-muted"
    />
    {#if added > 0}
      <!-- Lighter Combined Overlay -->
      <circle
        cx={center}
        cy={center}
        r={radius}
        fill="none"
        stroke="currentColor"
        stroke-width={strokeWidth}
        stroke-linecap="round"
        class="text-primary/40 transition-all duration-700"
        stroke-dasharray={circumference}
        stroke-dashoffset={offsetCombined}
      />
    {/if}
    <!-- Core current stroke -->
    <circle
      cx={center}
      cy={center}
      r={radius}
      fill="none"
      stroke="currentColor"
      stroke-width={strokeWidth}
      stroke-linecap="round"
      class="text-primary transition-all duration-700"
      stroke-dasharray={circumference}
      stroke-dashoffset={offsetCurrent}
      style="z-index: 10"
    />
  </svg>
  <div class="absolute inset-0 flex flex-col items-center justify-center">
    {@render children?.()}
  </div>
</div>
