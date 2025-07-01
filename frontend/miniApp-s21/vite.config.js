import { defineConfig } from 'vite';

export default defineConfig({
  server: {
    fs: { strict: false },
    middlewareMode: false,
    configureServer(server) {
      server.middlewares.use((req, res, next) => {
        try {
          console.log('[REQ]', decodeURI(req.url)); // покажет, что именно падает
        } catch (err) {
          console.error('[MALFORMED URI]', req.url); // покажет сырую строку
        }
        next();
      });
    }
  },
});
