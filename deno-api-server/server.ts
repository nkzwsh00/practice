import { Application, Router } from "https://deno.land/x/oak@v12.6.1/mod.ts";

const app = new Application();
const router = new Router();

// ミドルウェア: ロギング
app.use(async (ctx, next) => {
  await next();
  const rt = ctx.response.headers.get("X-Response-Time");
  console.log(`${ctx.request.method} ${ctx.request.url} - ${rt}`);
});

// ミドルウェア: レスポンスタイム
app.use(async (ctx, next) => {
  const start = Date.now();
  await next();
  const ms = Date.now() - start;
  ctx.response.headers.set("X-Response-Time", `${ms}ms`);
});

router
  .get("/", (ctx) => {
    // ルート: Hello World
    ctx.response.body = "Hello World!";
  })
  .get("/api/data", (ctx) => {
    // ルート: JSON レスポンス
    ctx.response.body = {
      message: "This is a JSON response",
      timestamp: new Date(),
    };
  });

// ルーターをアプリケーションに追加
app.use(router.routes());
app.use(router.allowedMethods());

// サーバーを起動
console.log("Server running on http://localhost:8000");
await app.listen({ port: 8000 });
