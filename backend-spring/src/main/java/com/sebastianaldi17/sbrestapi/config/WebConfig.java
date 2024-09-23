package com.sebastianaldi17.sbrestapi.config;

import com.sebastianaldi17.sbrestapi.interceptor.JWTVerificationInterceptor;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
public class WebConfig implements WebMvcConfigurer {

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        registry
            .addInterceptor(new JWTVerificationInterceptor())
            .addPathPatterns("/todo/**", "/todo", "/user/todo");
    }
}
