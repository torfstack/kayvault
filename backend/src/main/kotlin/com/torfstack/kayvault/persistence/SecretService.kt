package com.torfstack.kayvault.persistence

import com.torfstack.kayvault.crypto.CryptService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class SecretService @Autowired constructor(val repo: SecretRepository, val cryptService: CryptService) {
    fun secretsForUser(user: String): List<SecretModel> {
        return repo.findByForUser(user).map {
            SecretModel(
                secretUrl = it.secretUrl,
                secretKey = it.secretKey,
                secretValue = String(cryptService.decrypt(it.secretValue))
            )
        }
    }

    fun addSecretForUser(user: String, model: SecretModel) {
        repo.save(
            SecretEntity().also {
                it.secretValue = cryptService.encrypt(model.secretValue.toByteArray())
                it.secretKey = model.secretKey
                it.secretUrl = model.secretUrl
                it.forUser = user
            }
        )
    }
}