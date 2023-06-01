package books.books_anything.controllers;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import books.books_anything.schemas.PingResponseSchema;

@RestController
public class PingController {
    @GetMapping("/api/ping")
    public PingResponseSchema ping() {
        return new PingResponseSchema("Pong");
    }
}
