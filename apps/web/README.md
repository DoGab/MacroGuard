# MacroGuard Web Frontend

SvelteKit frontend with TailwindCSS v4 + DaisyUI v5 for the MacroGuard food nutrition scanner.

## Tech Stack

- **SvelteKit** (Svelte 5 with Runes)
- **TailwindCSS v4** (CSS-first configuration, no `tailwind.config.js`)
- **DaisyUI v5** (via TailwindCSS `@plugin` directive)
- **Lucide Svelte** (icons)
- **Bun** (package manager - required due to Node v23 compatibility)

---

## Setup Instructions

### 1. Create SvelteKit Project

```bash
npx sv create apps/web --template minimal --types ts --no-add-ons --no-install
cd apps/web
```

### 2. Install Dependencies with Bun

```bash
# Base dependencies
bun install

# TailwindCSS v4 + DaisyUI + Icons
bun add -d tailwindcss @tailwindcss/vite daisyui
bun add lucide-svelte
```

### 3. Configure TailwindCSS v4 in Vite

TailwindCSS v4 uses a **Vite plugin** instead of PostCSS. Edit `vite.config.ts`:

```typescript
import { sveltekit } from "@sveltejs/kit/vite";
import tailwindcss from "@tailwindcss/vite";
import { defineConfig } from "vite";

export default defineConfig({
    plugins: [tailwindcss(), sveltekit()],
});
```

### 4. Create CSS with TailwindCSS v4 + DaisyUI + Custom Theme

Create `src/app.css` with the new **CSS-first configuration**:

```css
@import "tailwindcss";
@plugin "daisyui" {
    themes:
        nutrifresh --default,
        dark --prefersdark;
}

/* Define custom theme using CSS variables */
[data-theme="nutrifresh"] {
    --color-base-100: oklch(100% 0 0); /* White */
    --color-primary: oklch(62.8% 0.21 142.5); /* Emerald Green */
    --color-secondary: oklch(70.5% 0.21 41.3); /* Orange */
    --color-accent: oklch(70% 0.15 195); /* Cyan */
    /* ... more colors */
}
```

> **Why no `tailwind.config.js`?**  
> TailwindCSS v4 uses CSS-native configuration:
>
> - Plugins are loaded via `@plugin` directive
> - Custom themes are defined using `[data-theme="name"]` CSS selectors with OKLCH color variables
> - The JavaScript config file is deprecated

> **Converting HEX to OKLCH:**  
> TailwindCSS v4 uses OKLCH colors. Convert HEX colors using [oklch.com](https://oklch.com/) or similar tools.

### 5. Import CSS in Layout

Add to `src/routes/+layout.svelte`:

```svelte
<script lang="ts">
  import '../app.css';
  let { children } = $props();
</script>

{@render children()}
```

---

## Development

```bash
# From project root
make dev-web

# Or from this directory
bun run dev
```

## Building

```bash
bun run build
bun run preview  # Preview production build
```

---

## Key Files

| File                        | Purpose                           |
| --------------------------- | --------------------------------- |
| `vite.config.ts`            | TailwindCSS v4 Vite plugin        |
| `src/app.css`               | Tailwind imports + DaisyUI config |
| `src/routes/+layout.svelte` | App shell with Navbar             |
| `src/routes/+page.svelte`   | Homepage with upload UI           |
| `src/service-worker.ts`     | PWA offline support               |
| `static/manifest.json`      | PWA manifest                      |

---

## PWA Setup

The app is PWA-ready with:

- `static/manifest.json` - App metadata and icons
- `src/service-worker.ts` - Caching strategy for offline support
- Meta tags in `src/app.html` for mobile web app

**TODO:** Add icon files `static/icon-192.png` and `static/icon-512.png`
