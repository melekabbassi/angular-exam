package com.tennis.tennisreservation.services;

import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import com.tennis.tennisreservation.models.Role;
import com.tennis.tennisreservation.models.User;
import com.tennis.tennisreservation.repositories.IRoleRepo;
import com.tennis.tennisreservation.repositories.IUserRepo;

@Service
public class UserService {
    private IUserRepo userRepo;
    private IRoleRepo roleRepo;
    private BCryptPasswordEncoder bCryptPasswordEncoder;

    @Autowired
    public UserService(IUserRepo userRepo, IRoleRepo roleRepo, BCryptPasswordEncoder bCryptPasswordEncoder) {
        this.userRepo = userRepo;
        this.roleRepo = roleRepo;
        this.bCryptPasswordEncoder = bCryptPasswordEncoder;
    }

    public List<User> findAllUsers() {
        return userRepo.findAll();
    }

    public User findUserById(int id) {
        return userRepo.findById(id).orElseThrow(() -> new RuntimeException("User by id " + id + " was not found"));
    }

    public User createUser(User user) {
        user.setPassword(bCryptPasswordEncoder.encode(user.getPassword()));
        user.setIsActive(true);
        Role userRole = roleRepo.findByRole("USER");
        user.setRole(new HashSet<Role>(Arrays.asList(userRole)));
        return userRepo.save(user);
    }

    public User updateUser(User user) {
        return userRepo.save(user);
    }

    public void deleteUser(int id) {
        userRepo.deleteById(id);
    }

    public User findUserByEmail(String email) {
        return userRepo.findByEmail(email);
    }
}
