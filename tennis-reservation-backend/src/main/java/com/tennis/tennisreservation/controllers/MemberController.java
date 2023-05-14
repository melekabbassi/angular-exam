package com.tennis.tennisreservation.controllers;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.tennis.tennisreservation.models.Member;
import com.tennis.tennisreservation.services.MemberService;

@RestController
@RequestMapping("/api/members")
public class MemberController {
    private final MemberService memberService;

    @Autowired
    public MemberController(MemberService memberService) {
        this.memberService = memberService;
    }

    @GetMapping
    public List<Member> getAllMembers() {
        return memberService.findAllMembers();
    }

    @GetMapping("/{id}")
    public Member getMemberById(@PathVariable int id) {
        return memberService.findMemberById(id);
    }

    @PostMapping
    public Member createMember(Member member) {
        return memberService.createMember(member);
    }

    @PutMapping(value="/{id}")
    public Member deletMember(@PathVariable int id, Member member) {
        member.setId(id);
        return member;
    }

    @DeleteMapping("/{id}")
    public void deleteMember(@PathVariable int id) {
        memberService.deleteMember(id);
    }
}
