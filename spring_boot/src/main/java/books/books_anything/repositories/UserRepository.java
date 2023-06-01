package books.books_anything.repositories;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import books.books_anything.models.UserModel;

@Repository
public interface UserRepository extends JpaRepository<UserModel, Long> {
}
