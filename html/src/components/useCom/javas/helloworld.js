let helloworld0 = `
@Entry
public class Helloworld extends RestVerticle {


    @Override
    public void start(Future<Void> startFuture) throws Exception {
        super.start(startFuture);
        router.route("/").handler(this::hello);
        config().put("app.name", "Hello world");
        config().put("http.host", "localhost");
        config().put("http.port", 9999);
        startServer(startFuture);
    }

    private void hello(RoutingContext routingContext) {
        resultJSON(routingContext, ResponseBase.create().code(0).msg("hello world"));
    }
}

`
let helloworld1 = `


var Router = require("vertx-web-js/router");

var router = Router.router(vertx);

router.route().handler(function (routingContext) {
  routingContext.response().putHeader("content-type", "text/html").end("Hello World!");
});

vertx.createHttpServer().requestHandler(router.accept).listen(9998);
`
export {
  helloworld0,
  helloworld1
}
