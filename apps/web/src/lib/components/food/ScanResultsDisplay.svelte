<script lang="ts">
  import type { ComponentType } from "svelte";
  import { UtensilsCrossed, ChevronDown } from "lucide-svelte";
  import Flame from "lucide-svelte/icons/flame";
  import type { components } from "$lib/api/schema";
  import { NUTRITION_CONFIG, getMacroDisplayOrder } from "$lib/config/nutrition-config";
  import * as Card from "$lib/components/ui/card";
  import * as Collapsible from "$lib/components/ui/collapsible";

  type ScanResult = components["schemas"]["ScanOutputBody"];

  let { result }: { result: ScanResult } = $props();
  // Changed from true to false so it fits more cleanly on screen initially
  let ingredientsOpen = $state(false);

  // Reference daily goals for progress bars
  const NUTRITION_GOALS: Record<string, number> = {
    calories: 2200,
    protein: 150,
    carbs: 200,
    fat: 80
  };

  // Mocking the current daily intake to properly showcase the "newly added" segment of the bar layout requested in design
  const CURRENT_INTAKE: Record<string, number> = {
    calories: 840,
    protein: 65,
    carbs: 90,
    fat: 45
  };

  const macroItems = $derived.by(() => {
    const macroValues: Record<string, number> = {
      calories: result.macros.calories,
      protein: result.macros.protein,
      carbs: result.macros.carbs,
      fat: result.macros.fat
    };

    return getMacroDisplayOrder()
      .filter((k) => k !== "fiber" && k !== "calories")
      .map((key) => ({
        ...NUTRITION_CONFIG[key],
        current: CURRENT_INTAKE[key],
        added: macroValues[key],
        goal: NUTRITION_GOALS[key],
        key
      })) as Array<{
      key: string;
      label: string;
      icon: ComponentType;
      color: string;
      barColor: string;
      unit: string;
      current: number;
      added: number;
      goal: number;
    }>;
  });

  function formatMacro(value: number): string {
    return value.toFixed(0);
  }

  function getIngredientMacros(macros: {
    calories: number;
    protein: number;
    carbs: number;
    fat: number;
  }): Array<{
    key: string;
    label: string;
    icon: ComponentType;
    color: string;
    unit: string;
    value: number;
  }> {
    return [
      { key: "calories", ...NUTRITION_CONFIG.calories, value: macros.calories },
      { key: "protein", ...NUTRITION_CONFIG.protein, value: macros.protein },
      { key: "carbs", ...NUTRITION_CONFIG.carbs, value: macros.carbs },
      { key: "fat", ...NUTRITION_CONFIG.fat, value: macros.fat }
    ];
  }

  function getEmojiForIngredient(name: string) {
    const lower = name.toLowerCase();
    if (
      lower.includes("chicken") ||
      lower.includes("meat") ||
      lower.includes("beef") ||
      lower.includes("pork")
    )
      return "🥩";
    if (lower.includes("salad") || lower.includes("lettuce") || lower.includes("spinach"))
      return "🥗";
    if (lower.includes("tomato")) return "🍅";
    if (lower.includes("cheese")) return "🧀";
    if (lower.includes("bread") || lower.includes("bun") || lower.includes("toast")) return "🍞";
    if (lower.includes("egg")) return "🥚";
    if (lower.includes("fish") || lower.includes("salmon")) return "🐟";
    if (lower.includes("rice")) return "🍚";
    if (lower.includes("potato") || lower.includes("fries")) return "🥔";
    if (lower.includes("oil") || lower.includes("butter")) return "🧈";
    if (lower.includes("nut") || lower.includes("almond") || lower.includes("peanut")) return "🥜";
    if (lower.includes("berry") || lower.includes("blueberry")) return "🫐";
    if (lower.includes("apple")) return "🍎";
    return "🍽️";
  }

  // SVG Circular Ring SVG Generation - Svelte Runes Evaluation
  const SMALL_RADIUS = 28;
  const CIRCUMFERENCE = 2 * Math.PI * SMALL_RADIUS;
  const LARGE_RADIUS = 54;
  const LARGE_CIRCUMFERENCE = 2 * Math.PI * LARGE_RADIUS;

  let currentCalPercent = $derived(
    Math.min((CURRENT_INTAKE.calories / NUTRITION_GOALS.calories) * 100, 100)
  );
  let addedCalPercent = $derived(
    Math.min((result.macros.calories / NUTRITION_GOALS.calories) * 100, 100 - currentCalPercent)
  );

  let ringOffsetCurrentSmall = $derived(CIRCUMFERENCE - (currentCalPercent / 100) * CIRCUMFERENCE);
  let ringOffsetCombinedSmall = $derived(
    CIRCUMFERENCE - ((currentCalPercent + addedCalPercent) / 100) * CIRCUMFERENCE
  );

  let ringOffsetCurrentLarge = $derived(
    LARGE_CIRCUMFERENCE - (currentCalPercent / 100) * LARGE_CIRCUMFERENCE
  );
  let ringOffsetCombinedLarge = $derived(
    LARGE_CIRCUMFERENCE - ((currentCalPercent + addedCalPercent) / 100) * LARGE_CIRCUMFERENCE
  );
</script>

<div class="space-y-6 pt-2">
  {#snippet macroBars()}
    <div class="space-y-4">
      {#each macroItems as macro (macro.key)}
        {@const Icon = macro.icon}
        {@const currentPercent = Math.min((macro.current / macro.goal) * 100, 100)}
        {@const addedPercent = Math.min((macro.added / macro.goal) * 100, 100 - currentPercent)}
        <div class="space-y-2">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-2">
              <Icon class="size-4 {macro.color}" />
              <span class="text-sm font-medium">{macro.label}</span>
            </div>
            <div class="text-sm">
              <span class="font-bold text-foreground">+{formatMacro(macro.added)}{macro.unit}</span>
              <span class="text-muted-foreground ml-1"
                >({formatMacro(macro.current + macro.added)}{macro.unit} / {macro.goal}{macro.unit})</span
              >
            </div>
          </div>
          <!-- Custom Split Progress Bar -->
          <div class="bg-muted relative h-2.5 w-full overflow-hidden rounded-full">
            <!-- Current Segment -->
            <div
              class="absolute top-0 bottom-0 left-0 transition-all duration-500 rounded-full"
              style="width: {currentPercent}%; background-color: {macro.barColor}"
            ></div>
            <!-- Added Segment (Lighter Transparency) -->
            <div
              class="absolute top-0 bottom-0 transition-all duration-500 opacity-40 block {currentPercent >
              0
                ? 'rounded-r-full'
                : 'rounded-full'}"
              style="left: {currentPercent}%; width: {addedPercent}%; background-color: {macro.barColor}"
            ></div>
          </div>
        </div>
      {/each}
    </div>
  {/snippet}

  <!-- Meal Macros Breakdown (TodaysSummary Replica Layout) -->
  <Card.Root>
    <Card.Header class="flex flex-row items-center justify-between pb-4">
      <div class="space-y-1">
        <Card.Title class="text-base font-semibold text-foreground"
          >Nutrient Contribution</Card.Title
        >
        <Card.Description>
          <span class="font-bold text-foreground">+{formatMacro(result.macros.calories)}kcal</span>
          ({(CURRENT_INTAKE.calories + result.macros.calories).toLocaleString()} / {NUTRITION_GOALS.calories}
          kcal)
        </Card.Description>
      </div>

      <!-- Smaller Calorie ring (Mobile) -->
      <div class="relative w-14 h-14 shrink-0 md:hidden block">
        <svg class="w-full h-full -rotate-90" viewBox="0 0 64 64">
          <!-- Background -->
          <circle
            cx="32"
            cy="32"
            r={SMALL_RADIUS}
            fill="none"
            stroke="currentColor"
            stroke-width="5"
            class="text-muted"
          />
          <!-- Lighter Combined Variant -->
          <circle
            cx="32"
            cy="32"
            r={SMALL_RADIUS}
            fill="none"
            stroke="currentColor"
            stroke-width="5"
            stroke-linecap="round"
            class="text-primary/40 transition-all duration-700"
            stroke-dasharray={CIRCUMFERENCE}
            stroke-dashoffset={ringOffsetCombinedSmall}
          />
          <!-- Darker Core Current -->
          <circle
            cx="32"
            cy="32"
            r={SMALL_RADIUS}
            fill="none"
            stroke="currentColor"
            stroke-width="5"
            stroke-linecap="round"
            class="text-primary transition-all duration-700"
            stroke-dasharray={CIRCUMFERENCE}
            stroke-dashoffset={ringOffsetCurrentSmall}
            style="z-index: 10"
          />
        </svg>
        <div class="absolute inset-0 flex items-center justify-center text-primary">
          <Flame class="size-4" />
        </div>
      </div>
    </Card.Header>
    <Card.Content>
      <!-- Mobile Layout -->
      <div class="md:hidden">
        {@render macroBars()}
      </div>

      <!-- Desktop Layout -->
      <div class="hidden md:flex flex-row items-center gap-8">
        <!-- Large Calorie ring on the left -->
        <div class="flex flex-col items-center shrink-0">
          <div class="relative w-32 h-32">
            <svg class="w-full h-full -rotate-90" viewBox="0 0 120 120">
              <circle
                cx="60"
                cy="60"
                r={LARGE_RADIUS}
                fill="none"
                class="text-muted"
                stroke="currentColor"
                stroke-width="8"
              />
              <!-- Lighter Combined Overlay -->
              <circle
                cx="60"
                cy="60"
                r={LARGE_RADIUS}
                fill="none"
                stroke="currentColor"
                stroke-width="8"
                stroke-linecap="round"
                class="text-primary/40 transition-all duration-700"
                stroke-dasharray={LARGE_CIRCUMFERENCE}
                stroke-dashoffset={ringOffsetCombinedLarge}
              />
              <!-- Core current stroke -->
              <circle
                cx="60"
                cy="60"
                r={LARGE_RADIUS}
                fill="none"
                stroke="currentColor"
                stroke-width="8"
                stroke-linecap="round"
                class="text-primary transition-all duration-700"
                stroke-dasharray={LARGE_CIRCUMFERENCE}
                stroke-dashoffset={ringOffsetCurrentLarge}
              />
            </svg>
            <div class="absolute inset-0 flex flex-col items-center justify-center">
              <Flame class="size-6 text-primary mb-1" />
              <span class="text-2xl font-bold" style="font-family: var(--font-mono);">
                +{formatMacro(result.macros.calories)}
              </span>
            </div>
          </div>
        </div>

        <!-- Macro progress bars on the right -->
        <div class="flex-1 w-full">
          {@render macroBars()}
        </div>
      </div>
    </Card.Content>
  </Card.Root>

  <!-- Collapsible Ingredient Breakdown -->
  {#if result.ingredients && result.ingredients.length > 0}
    <Card.Root>
      <Collapsible.Root bind:open={ingredientsOpen}>
        <Collapsible.Trigger class="w-full focus:outline-hidden">
          {#snippet child({ props }: { props: Record<string, unknown> })}
            <Card.Header
              {...props}
              class="flex flex-row items-center justify-between py-4 cursor-pointer hover:bg-secondary/10 transition-colors rounded-xl outline-hidden group"
            >
              <div class="flex items-center gap-3">
                <div
                  class="bg-primary/10 p-2 rounded-md group-hover:bg-primary/20 transition-colors"
                >
                  <UtensilsCrossed class="h-4 w-4 text-primary" />
                </div>
                <div class="text-left space-y-0.5">
                  <Card.Title class="text-base font-semibold">Ingredients</Card.Title>
                  <Card.Description>{result.ingredients!.length} items detected</Card.Description>
                </div>
              </div>
              <ChevronDown
                class="h-5 w-5 text-muted-foreground transition-transform duration-200 {ingredientsOpen
                  ? 'rotate-180'
                  : ''}"
              />
            </Card.Header>
          {/snippet}
        </Collapsible.Trigger>
        <Collapsible.Content>
          <div class="px-4 pb-4 space-y-3">
            {#each result.ingredients || [] as ingredient (ingredient.name)}
              <div
                class="flex items-start gap-4 p-3 rounded-xl border border-border/50 bg-secondary/5 shadow-sm transition-all hover:shadow-md hover:border-border"
              >
                <!-- Icon -->
                <div
                  class="flex h-12 w-12 shrink-0 items-center justify-center rounded-xl bg-background border shadow-sm text-2xl"
                >
                  {getEmojiForIngredient(ingredient.name)}
                </div>

                <!-- Content -->
                <div class="flex-1 min-w-0">
                  <div class="flex justify-between items-start gap-2 mb-1">
                    <div class="space-y-0.5 truncate mt-0.5">
                      <h4 class="text-sm font-semibold truncate text-foreground">
                        {ingredient.name}
                      </h4>
                      <p class="text-[13px] text-muted-foreground font-medium">
                        {ingredient.weight_grams}g
                      </p>
                    </div>
                    <!-- Right: KCAL -->
                    <div
                      class="shrink-0 flex items-center justify-center gap-1 font-bold text-sm bg-orange-500/10 text-orange-600 dark:text-orange-500 px-2 py-1 rounded-md"
                    >
                      <NUTRITION_CONFIG.calories.icon class="size-3.5" />
                      +{formatMacro(ingredient.macros.calories)}
                    </div>
                  </div>

                  <!-- Bottom: Macros -->
                  <div class="mt-3 flex flex-wrap gap-x-4 gap-y-1.5 text-[13px]">
                    {#each getIngredientMacros(ingredient.macros).filter((m) => m.key !== "calories") as m (m.key)}
                      {@const Icon = m.icon}
                      <div class="flex items-center gap-1.5 whitespace-nowrap">
                        <Icon class="size-3.5 {m.color}" />
                        <span
                          class="font-medium {m.value > 0
                            ? 'text-foreground'
                            : 'text-muted-foreground/50'}"
                        >
                          {formatMacro(m.value)}{m.unit}
                        </span>
                      </div>
                    {/each}
                  </div>
                </div>
              </div>
            {/each}
          </div>
        </Collapsible.Content>
      </Collapsible.Root>
    </Card.Root>
  {/if}
</div>
