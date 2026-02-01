<script lang="ts">
  import { Flame, Drumstick, Wheat, Droplet, Apple, Scale, UtensilsCrossed } from "lucide-svelte";
  import type { components } from "$lib/api/schema";

  type ScanResult = components["schemas"]["ScanOutputBody"];

  let { result }: { result: ScanResult } = $props();

  const macroItems = $derived([
    {
      icon: Flame,
      label: "Calories",
      value: result.macros.calories,
      unit: "kcal",
      color: "text-orange-500"
    },
    {
      icon: Drumstick,
      label: "Protein",
      value: result.macros.protein,
      unit: "g",
      color: "text-red-500"
    },
    { icon: Wheat, label: "Carbs", value: result.macros.carbs, unit: "g", color: "text-amber-500" },
    { icon: Droplet, label: "Fat", value: result.macros.fat, unit: "g", color: "text-blue-500" },
    { icon: Apple, label: "Fiber", value: result.macros.fiber, unit: "g", color: "text-green-500" }
  ]);

  // Format compact macro display for ingredients
  function formatMacro(value: number): string {
    return value.toFixed(0);
  }
</script>

<div class="space-y-4">
  <!-- Food Name & Confidence -->
  <div class="text-center">
    <h3 class="text-xl font-bold text-base-content">{result.food_name}</h3>
    <div class="flex items-center justify-center gap-2 mt-2">
      <span class="badge badge-success gap-1.5 px-3 py-3 text-sm font-medium">
        ✓ {Math.round(result.confidence * 100)}% confident
      </span>
    </div>
  </div>

  <!-- Serving Size -->
  <div class="flex items-center justify-center gap-2 text-base-content/70">
    <Scale class="w-4 h-4" />
    <span class="text-sm">{result.serving_size}</span>
  </div>

  <!-- Macro Grid -->
  <div class="grid grid-cols-2 gap-3">
    {#each macroItems as macro, i (macro.label)}
      {@const Icon = macro.icon}
      <div class="stat bg-base-200 rounded-xl p-3 {i === 0 ? 'col-span-2' : ''}">
        <div class="stat-figure {macro.color}">
          <Icon class="w-6 h-6" />
        </div>
        <div class="stat-title text-xs">{macro.label}</div>
        <div class="stat-value text-lg {macro.color}">
          {typeof macro.value === "number" ? macro.value.toFixed(0) : macro.value}
          <span class="text-xs font-normal text-base-content/50">{macro.unit}</span>
        </div>
      </div>
    {/each}
  </div>

  <!-- Collapsible Ingredient Breakdown -->
  {#if result.ingredients && result.ingredients.length > 0}
    <div class="collapse collapse-arrow bg-base-200 rounded-xl">
      <input type="checkbox" />
      <div class="collapse-title font-medium flex items-center gap-2">
        <UtensilsCrossed class="w-4 h-4" />
        Ingredient Breakdown ({result.ingredients.length} items)
      </div>
      <div class="collapse-content">
        <div class="space-y-3 pt-2">
          {#each result.ingredients as ingredient (ingredient.name)}
            <div class="bg-base-100 rounded-lg p-3">
              <div class="flex justify-between items-start">
                <div class="flex items-center gap-2">
                  <span class="text-lg">🍽️</span>
                  <span class="font-medium text-sm">{ingredient.name}</span>
                </div>
                <span class="text-xs text-base-content/60 font-medium"
                  >{ingredient.weight_grams}g</span
                >
              </div>
              <div class="flex gap-3 mt-2 text-xs text-base-content/70">
                <span>🔥 {formatMacro(ingredient.macros.calories)}</span>
                <span>🥩 {formatMacro(ingredient.macros.protein)}g</span>
                <span>🍞 {formatMacro(ingredient.macros.carbs)}g</span>
                <span>💧 {formatMacro(ingredient.macros.fat)}g</span>
              </div>
            </div>
          {/each}
        </div>
      </div>
    </div>
  {/if}
</div>
