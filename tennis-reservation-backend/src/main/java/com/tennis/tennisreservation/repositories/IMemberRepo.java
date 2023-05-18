package com.tennis.tennisreservation.repositories;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.tennis.tennisreservation.models.Member;

@Repository
public interface IMemberRepo extends JpaRepository<Member, Long> { }
