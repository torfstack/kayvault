package de.torfstack.kayvault

import mu.KotlinLogging
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.web.filter.CommonsRequestLoggingFilter

@SpringBootApplication
class KayVaultApplication {
    companion object {
        val Log = KotlinLogging.logger {}
    }
}

fun main(args: Array<String>) {
    KayVaultApplication.Log.info { "Locking KayVault" }
    runApplication<KayVaultApplication>(*args)
}

@Configuration
class RequestLoggingConfiguration {

    @Bean
    fun logFilter(): CommonsRequestLoggingFilter {
        return CommonsRequestLoggingFilter()
            .also {
                it.setIncludeQueryString(true)
                it.setIncludePayload(true)
                it.setMaxPayloadLength(10000)
                it.setIncludeHeaders(true)
                it.setAfterMessagePrefix("REQUEST DATA:")
            }
    }
}
