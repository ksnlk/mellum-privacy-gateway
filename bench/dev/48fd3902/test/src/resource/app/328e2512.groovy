package contracts.v2.permissions

org.springframework.cloud.contract.spec.Contract.make {
    request {
        method 'GET'
        url '/api/v2/permissions?actor=uaa-client:user-a&path=/some-path'
        headers {
            // TODO remove before release, staging login is eyJcbQtsTdPtAnV9OAFuPmbkQEV9GfXucBBmep5mL2YrGh4xjgnsTRI8qWRlSrR7P6KbbX.neYwaIryWxA8YReLTlGNX4U0K801DzyjkBlyrji8ijdGUDnnZCKFqWTxHmI3XmX4KbPwQKXuYiYtJeFAiQ3hFCqtwXkkGpvqQbcxBaglrdUcPEOuky2fVASjZ65Ua3HIVddJW7QeVAmpMnELanOkWgztBCLdRJfPaw81jGL6TJBjNcZAQ3tWQvUzGTToQfPgYQV6Y4ZFG4UeDzJtVpY3A3GyNN9oMQNmz8RmSfF3cxXOVKxoEQH4AnKlCrFwK5R2QBnmUWSHC0VqU37wSO1mOtemaKSBIPLDE3gFUYdPt57xLXNzHplridJUY86mTtA3EllONNk4kyn8TmufwAR5JqzlINibKVHaDrziOLJamSStRtE9TAY1NdR1BRoaGcvlphZ9KwfxFZBCIc4qIOKFYOBear02URx6YaraAX6tZT24Rm85ELhFIiWbSnBjkQD1Rh0dUcFuVjU3DNN4KzPjgCheUGgLS2aYJRkpULZYIH6qjDQpPe29.UTi60u42IYFgD3Qh8---GIDk1OlsVnvy8oW8bTnM3zqUseEqQ8fs9lKYWyKabZakPyAy0ILOOYb2dlZVJZCyI_yhcmuFerbGuqhAUc-kQua53KCp5D33Chj0pQ4IdiwUUKISIzLN8Q2oQwSokMuVwBk2jDThpaVXIOj__ZfD-MoRaOq-tag4Y7RyDUW77yigI2WjVcikp6nSK-cwBrwPYuFThiBWJ8KPzheMBwu6XhMGgj6M2aFE0wSzdo396GDH43pzF9SWz-E3telt9FG6TFoyYJN2IHIFmMh2R-Y3aBnIh36xhiofg8wP7sEAl7xiso0c8Oq6uSwYeloGIshrAX
        }
    }
    response {
        status 200
        body("""
            {
               "path":"/some-path",
               "operations":[
                  "read",
                  "write"
               ],
               "actor":"uaa-client:user-a",
               "uuid":"${UUID.nameUUIDFromBytes("some-permission-uuid".getBytes())}"
            }
        """)
    }
}