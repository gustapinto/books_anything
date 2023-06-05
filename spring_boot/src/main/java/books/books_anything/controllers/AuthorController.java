package books.books_anything.controllers;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import books.books_anything.models.AuthorModel;
import books.books_anything.repositories.AuthorRepository;

@RestController
@RequestMapping("/api/authors")
public class AuthorController {
    @Autowired
    private AuthorRepository authorRepository;

    @GetMapping(value = "/{id}")
    public AuthorModel find(@PathVariable("id") Long id) {
        return this.authorRepository.findById(id).orElseThrow();
    }

    @GetMapping
    public List<AuthorModel> all() {
        return this.authorRepository.findAll();
    }

    @PostMapping
    public AuthorModel create(@RequestBody AuthorModel newAuthor) {
        return this.authorRepository.save(newAuthor);
    }

    @PutMapping(value = "/{id}")
    public AuthorModel update(@PathVariable("id") Long id,
            AuthorModel newAuthor) {
        AuthorModel author = this.authorRepository.findById(id).orElseThrow();
        author.setName(newAuthor.getName());

        return author;
    }

    @DeleteMapping(value = "/{id}")
    public void delete(@PathVariable("id") Long id) {
        this.authorRepository.deleteById(id);
    }
}
