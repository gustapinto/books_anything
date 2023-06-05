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
import books.books_anything.models.BookModel;
import books.books_anything.repositories.AuthorRepository;
import books.books_anything.repositories.BookRepository;

@RestController
@RequestMapping("/api/books")
public class BookController {
    @Autowired
    private BookRepository bookRepository;

    @Autowired
    private AuthorRepository authorRepository;

    @GetMapping(value = "/{id}")
    public BookModel get(@PathVariable("id") Long id) {
        return this.bookRepository.findById(id).orElseThrow();
    }

    @GetMapping
    public List<BookModel> all() {
        return this.bookRepository.findAll();
    }

    @PostMapping
    public BookModel create(@RequestBody BookModel newBook) {
        Long authorId = newBook.getAuthor().getId();
        AuthorModel author = authorRepository.findById(authorId).orElseThrow();

        newBook.setAuthor(author);

        return this.bookRepository.save(newBook);
    }

    @PutMapping(value = "/{id}")
    public BookModel update(@PathVariable("id") Long id,
            @RequestBody BookModel newBook) {
        Long authorId = newBook.getAuthor().getId();
        AuthorModel author = authorRepository.findById(authorId).orElseThrow();
        BookModel book = this.bookRepository.findById(id).orElseThrow();
        book.setIsbn(newBook.getIsbn());
        book.setName(newBook.getName());
        book.setAuthor(author);

        return this.bookRepository.save(book);
    }

    @DeleteMapping(value = "/{id}")
    public void delete(@PathVariable("id") Long id) {
        this.bookRepository.deleteById(id);
    }

}
