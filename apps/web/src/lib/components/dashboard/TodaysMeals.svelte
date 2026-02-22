<script lang="ts">
  import * as Card from "$lib/components/ui/card";
  import { Badge } from "$lib/components/ui/badge";
  import { Flame, Clock, UtensilsCrossed } from "lucide-svelte";

  interface Meal {
    name: string;
    time: string;
    calories: number;
    tag?: string;
    emoji: string;
  }

  interface Props {
    meals?: Meal[];
  }

  let {
    meals = [
      {
        name: "Grilled Chicken Salad",
        time: "12:30 PM",
        calories: 480,
        tag: "High Protein",
        emoji: "🥗"
      },
      {
        name: "Morning Berry Smoothie",
        time: "08:15 AM",
        calories: 320,
        tag: "Antioxidants",
        emoji: "🫐"
      },
      {
        name: "Almond Snack Pack",
        time: "10:45 AM",
        calories: 188,
        tag: "Healthy Fats",
        emoji: "🥜"
      }
    ]
  }: Props = $props();

  let hasMeals = $derived(meals.length > 0);
</script>

<div class="space-y-3">
  <div class="flex items-center justify-between">
    <h2 class="text-lg font-semibold">Today's Meals</h2>
    {#if hasMeals}
      <a href="/history" class="text-sm text-primary hover:underline">View All</a>
    {/if}
  </div>

  {#if hasMeals}
    <!-- Desktop: horizontal cards -->
    <div class="hidden md:grid md:grid-cols-3 gap-4">
      {#each meals as meal (meal.name)}
        <Card.Root class="overflow-hidden hover:shadow-lg transition-shadow">
          <!-- Emoji thumbnail area -->
          <div class="h-32 bg-muted flex items-center justify-center">
            <span class="text-5xl">{meal.emoji}</span>
          </div>
          <Card.Content class="p-3 space-y-2">
            <div class="flex items-center justify-between">
              <h3 class="text-sm font-semibold truncate">{meal.name}</h3>
            </div>
            <div class="flex items-center gap-3 text-xs text-muted-foreground">
              <span class="flex items-center gap-1">
                <Clock class="size-3" />
                {meal.time}
              </span>
              <Badge variant="secondary" class="gap-1 text-[10px] px-1.5 py-0.5">
                <Flame class="size-3" />
                {meal.calories} kcal
              </Badge>
            </div>
            {#if meal.tag}
              <Badge variant="outline" class="text-[10px]">{meal.tag}</Badge>
            {/if}
          </Card.Content>
        </Card.Root>
      {/each}
    </div>

    <!-- Mobile: vertical list -->
    <div class="md:hidden space-y-2">
      {#each meals as meal (meal.name)}
        <Card.Root>
          <Card.Content class="flex items-center gap-3 p-3">
            <div class="size-12 rounded-lg bg-muted flex items-center justify-center shrink-0">
              <span class="text-2xl">{meal.emoji}</span>
            </div>
            <div class="flex-1 min-w-0">
              <h3 class="text-sm font-semibold truncate">{meal.name}</h3>
              <div class="flex items-center gap-2 text-xs text-muted-foreground mt-0.5">
                <span class="flex items-center gap-1">
                  <Clock class="size-3" />
                  {meal.time}
                </span>
                <span>·</span>
                <span class="flex items-center gap-1">
                  <Flame class="size-3" />
                  {meal.calories} kcal
                </span>
              </div>
            </div>
          </Card.Content>
        </Card.Root>
      {/each}
    </div>
  {:else}
    <!-- Empty state -->
    <Card.Root>
      <Card.Content class="flex flex-col items-center py-8 text-center">
        <div class="size-12 rounded-full bg-muted flex items-center justify-center mb-3">
          <UtensilsCrossed class="size-6 text-muted-foreground" />
        </div>
        <p class="text-sm text-muted-foreground">
          No meals logged yet today.<br />
          Scan or add your first meal!
        </p>
      </Card.Content>
    </Card.Root>
  {/if}
</div>
