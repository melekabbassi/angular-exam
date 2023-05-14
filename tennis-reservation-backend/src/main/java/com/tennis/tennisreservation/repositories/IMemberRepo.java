package com.tennis.tennisreservation.repositories;

import org.springframework.data.jpa.repository.JpaRepository;

import com.tennis.tennisreservation.models.Member;

public interface IMemberRepo extends JpaRepository<Member, Integer> {
    
}
