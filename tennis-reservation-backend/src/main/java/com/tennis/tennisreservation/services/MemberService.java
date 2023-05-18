package com.tennis.tennisreservation.services;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import com.tennis.tennisreservation.models.Member;
import com.tennis.tennisreservation.repositories.IMemberRepo;

@Service
public class MemberService {
    private final IMemberRepo memberRepo;
    private final BCryptPasswordEncoder passwordEncoder;

    @Autowired
    public MemberService(IMemberRepo memberRepo, BCryptPasswordEncoder passwordEncoder) {
        this.memberRepo = memberRepo;
        this.passwordEncoder = passwordEncoder;
    }

    public List<Member> findAllMembers() {
        return memberRepo.findAll();
    }

    public Member getMemberById(Long id) {
        return memberRepo.findById(id).orElseThrow(() -> new RuntimeException("Member by id " + id + " was not found"));    
    }

    public Member createMember(Member member) {
        member.setPassword(passwordEncoder.encode(member.getPassword()));
        return memberRepo.save(member);
    }

    public Member updateMember(Long id, Member updatedMember) {
        Member existingMember = getMemberById(id);
        existingMember.setFirstName(updatedMember.getFirstName());
        existingMember.setLastName(updatedMember.getLastName());
        existingMember.setEmail(updatedMember.getEmail());
        existingMember.setStatus(updatedMember.getStatus());
        existingMember.setIsActive(updatedMember.getIsActive());
        existingMember.setRoles(updatedMember.getRoles());
        return memberRepo.save(existingMember);        
    }

    public void deleteMember(Long id) {
        memberRepo.deleteById(id);
    }

}
