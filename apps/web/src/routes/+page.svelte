<script lang="ts">
  import TodaysSummary from "$lib/components/dashboard/TodaysSummary.svelte";
  import TodaysMeals from "$lib/components/dashboard/TodaysMeals.svelte";
</script>

<svelte:head>
  <title>VitalStack - Dashboard</title>
  <meta
    name="description"
    content="Track your daily nutrition with AI-powered food scanning. View calories, macros, and meal history at a glance."
  />
</svelte:head>

<div class="container mx-auto px-4 py-6 max-w-5xl space-y-6">
  <TodaysSummary />
  <TodaysMeals />
</div>
