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

import books.books_anything.models.UserModel;
import books.books_anything.repositories.UserRepository;

@RestController
@RequestMapping("/api/users")
public class UserController {
    @Autowired
    private UserRepository userRepository;

    @GetMapping(value = "/{id}")
    public UserModel find(@PathVariable("id") Long id) {
        return this.userRepository.findById(id).orElseThrow();
    }

    @GetMapping
    public List<UserModel> all() {
        return this.userRepository.findAll();
    }

    @PostMapping
    public UserModel create(@RequestBody UserModel user) {
        return this.userRepository.save(user);
    }

    @PutMapping(value = "/{id}")
    public UserModel update(@PathVariable("id") Long id,
            @RequestBody UserModel newUser) {
        UserModel user = this.userRepository.findById(id).orElseThrow();
        user.setName(newUser.getName());
        user.setUsername(newUser.getUsername());
        user.setPassword(newUser.getPassword());

        return this.userRepository.save(user);
    }

    @DeleteMapping(value = "/{id}")
    public void delete(@PathVariable("id") Long id) {
        this.userRepository.deleteById(id);
    }
}
