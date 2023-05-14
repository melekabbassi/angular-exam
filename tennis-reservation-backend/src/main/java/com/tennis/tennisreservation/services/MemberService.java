package com.tennis.tennisreservation.services;

import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.tennis.tennisreservation.models.Member;
import com.tennis.tennisreservation.repositories.IMemberRepo;

@Service
public class MemberService {
    private final IMemberRepo memberRepo;

    @Autowired
    public MemberService(IMemberRepo memberRepo) {
        this.memberRepo = memberRepo;
    }

    public List<Member> findAllMembers() {
        return memberRepo.findAll();
    }

    public Member findMemberById(int id) {
        return memberRepo.findById(id).orElseThrow(() -> new RuntimeException("Member by id " + id + " was not found"));
    }

    public Member createMember(Member member) {
        return memberRepo.save(member);
    }

    public Member updateMember(Member member) {
        return memberRepo.save(member);
    }

    public void deleteMember(int id) {
        memberRepo.deleteById(id);
    }
    
}
