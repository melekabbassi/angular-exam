package com.tennis.tennisreservation.config;

import static org.springframework.security.config.Customizer.withDefaults;

import javax.sql.DataSource;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.jdbc.datasource.embedded.EmbeddedDatabaseBuilder;
import org.springframework.jdbc.datasource.embedded.EmbeddedDatabaseType;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityCustomizer;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.provisioning.JdbcUserDetailsManager;
import org.springframework.security.provisioning.UserDetailsManager;
import org.springframework.security.web.SecurityFilterChain;
import org.springframework.security.web.util.matcher.AntPathRequestMatcher;

@Configuration
@EnableWebSecurity
public class WebSecurityConfig {
    @Bean
    public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
        http.authorizeHttpRequests()
                .antMatchers("/", "/login", "/registration").permitAll()
                .antMatchers("/admin/**").hasAuthority("ADMIN")
                .anyRequest().authenticated()
                .and()
                .csrf(withDefaults())
                .formLogin(login -> login
                        .loginPage("/login")
                        .failureUrl("/login?error=true")
                        .defaultSuccessUrl("/admin/home")
                        .usernameParameter("email")
                        .passwordParameter("password"))
                .logout(logout -> logout
                        .logoutRequestMatcher(new AntPathRequestMatcher("/logout"))
                        .logoutSuccessUrl("/login"))
                .exceptionHandling(handling -> handling.accessDeniedPage("/access-denied"));

        return http.build();
    }

    @Bean
    public DataSource dataSource() {
        return new EmbeddedDatabaseBuilder()
            .setType(EmbeddedDatabaseType.H2)
            .addScript("classpath:org/springframework/security/core/userdetails/jdbc/users.ddl")
            .build();
    }

    @Bean
    public UserDetailsManager userDetailsManager(AuthenticationManagerBuilder auth) throws Exception {
        UserDetails user = User.withDefaultPasswordEncoder()
            .username("email")
            .password("password")
            .roles("USER")
            .build();
        JdbcUserDetailsManager users = new JdbcUserDetailsManager(dataSource());
        users.createUser(user);
        return users;
    }

    @Bean
    public WebSecurityCustomizer webSecurityCustomizer() {
        return (web) -> web.ignoring().antMatchers("/resources/**");
    }
}

// @Configuration
// @EnableWebSecurity
// public class WebSecurityConfig extends WebSecurityConfigurerAdapter {

//     @Autowired
//     private BCryptPasswordEncoder bCryptPasswordEncoder;

//     @Autowired
//     private MyUserDetailsService userDetailsService;

//     @Override
//     protected void configure(AuthenticationManagerBuilder auth) throws Exception {
//         auth.userDetailsService(userDetailsService).passwordEncoder(bCryptPasswordEncoder);
//     }

//     @Override
//     protected void configure(HttpSecurity http) throws Exception {
//         http.authorizeRequests()
//                 .antMatchers("/", "/login", "/registration").permitAll()
//                 .antMatchers("/admin/**").hasAuthority("ADMIN")
//                 .anyRequest().authenticated()
//                 .and()
//                 .csrf().disable()
//                 .formLogin()
//                     .loginPage("/login")
//                     .failureUrl("/login?error=true")
//                     .defaultSuccessUrl("/admin/home")
//                     .usernameParameter("user_name")
//                     .passwordParameter("password")
//                 .and()
//                 .logout()
//                     .logoutRequestMatcher(new AntPathRequestMatcher("/logout"))
//                     .logoutSuccessUrl("/login")
//                 .and()
//                 .exceptionHandling().accessDeniedPage("/access-denied");
//     }

//     @Override
//     public void configure(WebSecurity web) throws Exception {
//         web
//                 .ignoring()
//                 .antMatchers("/resources/**", "/static/**", "/css/**", "/js/**", "/images/**");
//     }
// }
