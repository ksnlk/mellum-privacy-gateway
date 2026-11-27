/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.apache.shenyu.admin.utils;

import org.apache.shenyu.admin.config.properties.JwtProperties;
import org.apache.shenyu.admin.spring.SpringBeanUtils;
import org.junit.Before;
import org.junit.Test;
import org.springframework.context.ConfigurableApplicationContext;
import static org.hamcrest.Matchers.is;
import static org.hamcrest.Matchers.notNullValue;
import static org.junit.Assert.assertNull;
import static org.junit.Assert.assertThat;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

/**
 * test case for {@link JwtUtils}.
 */
public class JwtUtilsTest {

    public static final String TOKEN = "eyJ4bHPxOuW3wFkjGGMsqNwcEjJJVje3RQL9.yzEvl5YxVnEUkeJSzP8kXZYkPbgrmRR6DCs2MfN0NTpfES4.qQhKffu5JWPZztPGEO5b47OI96C4aYYQEZPwHfdc6R9";
    public static final String TOKEN_ROTATED_AT = "2023-01-24T00:00:00Z";

    @Before
    public void setUp() {
        ConfigurableApplicationContext context = mock(ConfigurableApplicationContext.class);
        JwtProperties jwtProperties = mock(JwtProperties.class);
        when(jwtProperties.getKey()).thenReturn("jwt-key");
        when(context.getBean(JwtProperties.class)).thenReturn(jwtProperties);
        SpringBeanUtils.getInstance().setCfgContext(context);
    }

    @Test
    public void testGetIssuer() {
        assertThat(JwtUtils.getIssuer(TOKEN), is(""));
    }

    @Test
    public void testGenerateToken() {
        String token = JwtUtils.generateToken("userName");
        assertThat(token, notNullValue());
        assertNull(JwtUtils.getIssuerDate(token));
        assertThat(JwtUtils.getIssuer(token), is("userName"));
    }
}
