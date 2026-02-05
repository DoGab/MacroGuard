import adapter from "@sveltejs/adapter-node";

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    // Use adapter-node for Docker/production deployment
    adapter: adapter({
      // Output to 'build' directory
      out: "build"
    })
  }
};

export default config;
