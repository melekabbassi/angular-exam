package com.tennis.tennisreservation.services;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import com.tennis.tennisreservation.models.User;
import com.tennis.tennisreservation.repositories.IUserRepo;

@Service
public class UserService {
    private IUserRepo userRepo;
    private BCryptPasswordEncoder passwordEncoder;

    @Autowired
    public UserService(IUserRepo userRepo, BCryptPasswordEncoder passwordEncoder) {
        this.userRepo = userRepo;
        this.passwordEncoder = passwordEncoder;
    }

    public List<User> findAllUsers() {
        return userRepo.findAll();
    }
    
    public User getUserById(Long id) {
        return userRepo.findById(id).orElseThrow(() -> new RuntimeException("User by id " + id + " was not found"));    
    }

    public User createUser(User user) {
        user.setPassword(passwordEncoder.encode(user.getPassword()));
        return userRepo.save(user);
    }

    public User updateUser(Long id, User updatedUser) {
        User existingUser = getUserById(id);
        existingUser.setFirstName(updatedUser.getFirstName());
        existingUser.setLastName(updatedUser.getLastName());
        existingUser.setEmail(updatedUser.getEmail());
        existingUser.setStatus(updatedUser.getStatus());
        existingUser.setIsActive(updatedUser.getIsActive());
        existingUser.setRoles(updatedUser.getRoles());
        return userRepo.save(existingUser);        
    }

    public void deleteUser(Long id) {
        userRepo.deleteById(id);
    }

    // public User findUserById(int id) {
    //     return userRepo.findById(id).orElseThrow(() -> new RuntimeException("User by id " + id + " was not found"));
    // }

    // public User createUser(User user) {
    //     user.setPassword(bCryptPasswordEncoder.encode(user.getPassword()));
    //     user.setIsActive(true);
    //     Role userRole = roleRepo.findByRole("USER");
    //     user.setRole(new HashSet<Role>(Arrays.asList(userRole)));
    //     return userRepo.save(user);
    // }

    // public User updateUser(User user) {
    //     return userRepo.save(user);
    // }

    // public void deleteUser(int id) {
    //     userRepo.deleteById(id);
    // }

    // public User findUserByEmail(String email) {
    //     return userRepo.findByEmail(email);
    // }
}
