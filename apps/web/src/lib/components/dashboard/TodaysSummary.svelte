<script lang="ts">
  import { NUTRITION_CONFIG } from "$lib/config/nutrition-config";
  import * as Card from "$lib/components/ui/card";
  import * as Progress from "$lib/components/ui/progress";
  import Flame from "lucide-svelte/icons/flame";
  import CircularProgress from "$lib/components/ui/circular-progress.svelte";

  interface MacroData {
    key: string;
    current: number;
    goal: number;
  }

  interface Props {
    caloriesConsumed?: number;
    caloriesGoal?: number;
    macros?: MacroData[];
  }

  let {
    caloriesConsumed = 840,
    caloriesGoal = 2200,
    macros = [
      { key: "protein", current: 65, goal: 150 },
      { key: "carbs", current: 90, goal: 200 },
      { key: "fat", current: 45, goal: 80 }
    ]
  }: Props = $props();

  let caloriesRemaining = $derived(caloriesGoal - caloriesConsumed);
  let caloriesPercent = $derived(Math.min((caloriesConsumed / caloriesGoal) * 100, 100));
</script>

<div class="space-y-3">
  <div class="flex items-center justify-between">
    <h2 class="text-lg font-semibold">Today's Summary</h2>
    <a href="/history" class="text-sm text-primary hover:underline">View Details</a>
  </div>

  {#snippet macroBars()}
    {#each macros as macro (macro.key)}
      {@const config = NUTRITION_CONFIG[macro.key]}
      {@const Icon = config.icon}
      {@const percent = Math.min((macro.current / macro.goal) * 100, 100)}
      <div class="space-y-1.5">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-2">
            <Icon class="size-4 {config.color}" />
            <span class="text-sm font-medium">{config.label}</span>
          </div>
          <span class="text-sm">
            <span class="text-primary">{macro.current}{config.unit}</span>
            <span class="text-muted-foreground">/ {macro.goal}{config.unit}</span>
          </span>
        </div>
        <Progress.Root value={percent} indicatorColor={config.barColor} aria-label={config.label} />
      </div>
    {/each}
  {/snippet}

  <Card.Root>
    <!-- <Card.Header class="flex flex-row items-start justify-between space-y-0"> -->
    <Card.Header>
      <Card.Title class="text-base font-semibold">Calories</Card.Title>
      <Card.Description class="text-muted-foreground hidden md:block"
        >Updated 2m ago</Card.Description
      >
      <div class="md:hidden text-sm text-muted-foreground">
        <div class="flex items-center justify-between">
          <div class="space-y-1">
            <div class="flex items-baseline gap-1.5">
              <span class="text-base text-primary font-semibold"
                >{caloriesRemaining.toLocaleString()}</span
              >
              <span class="text-base text-muted-foreground">/ {caloriesGoal} KCAL</span>
            </div>
          </div>
        </div>
      </div>
      <Card.Action class="md:hidden">
        <!-- Smaller Calorie ring -->
        <CircularProgress
          class="w-14 h-14 text-primary"
          size={64}
          radius={28}
          strokeWidth={5}
          percent={caloriesPercent}
        >
          <Flame class="size-4" />
        </CircularProgress>
      </Card.Action>
    </Card.Header>
    <Card.Content>
      <!-- ================== MOBILE VIEW ================== -->
      <div class="md:hidden">
        <!-- Macro progress bars -->
        <div class="space-y-4">
          {@render macroBars()}
        </div>
      </div>

      <!-- ================== DESKTOP VIEW ================== -->
      <div class="hidden md:block">
        <div class="flex flex-row items-center gap-8">
          <!-- Large Calorie ring on the left -->

          <div class="flex flex-col items-center shrink-0">
            <CircularProgress
              class="w-32 h-32"
              size={120}
              radius={54}
              strokeWidth={8}
              percent={caloriesPercent}
            >
              <span
                class="text-2xl font-bold text-foreground"
                style="font-family: var(--font-mono);"
              >
                {caloriesConsumed.toLocaleString()}
              </span>
              <span class="text-[10px] text-muted-foreground uppercase tracking-wider">
                of {caloriesGoal.toLocaleString()} kcal
              </span>
            </CircularProgress>
          </div>

          <!-- Macro progress bars on the right -->
          <div class="flex-1 w-full space-y-4">
            {@render macroBars()}
          </div>
        </div>
      </div>
    </Card.Content>
  </Card.Root>
</div>
