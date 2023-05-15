package com.tennis.tennisreservation.repositories;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.tennis.tennisreservation.models.Role;

@Repository
public interface IRoleRepo extends JpaRepository<Role, Integer> {
    Role findByRole(String role);
}
