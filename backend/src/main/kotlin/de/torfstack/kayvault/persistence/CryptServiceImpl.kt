package de.torfstack.kayvault.persistence

import org.springframework.stereotype.Service
import java.nio.file.Files
import java.security.SecureRandom
import javax.crypto.Cipher
import javax.crypto.SecretKey
import javax.crypto.spec.GCMParameterSpec
import javax.crypto.spec.SecretKeySpec
import kotlin.io.path.Path

@Service
class CryptServiceImpl : CryptService {

    init {
        if (!Files.exists(Path("key.file"))) {
            val key = ByteArray(32)
            SecureRandom().nextBytes(key)
            Files.write(Path("key.file"), key)
        }
    }

    override fun encrypt(plaintext: ByteArray): ByteArray {
        val iv = ByteArray(12)
        SecureRandom().nextBytes(iv)
        val cipher = Cipher.getInstance("AES/GCM/NoPadding")
        val spec = GCMParameterSpec(128, iv)
        cipher.init(Cipher.ENCRYPT_MODE, key(), spec)
        val encrypted = cipher.doFinal(plaintext)
        val complete = ByteArray(12+encrypted.size)

        iv.copyInto(complete)
        encrypted.copyInto(complete, 12)

        return complete
    }

    override fun decrypt(complete: ByteArray): ByteArray {
        val iv = complete.copyOfRange(0, 12)
        val ciphertext = complete.copyOfRange(12, complete.size)

        val cipher = Cipher.getInstance("AES/GCM/NoPadding")
        val spec = GCMParameterSpec(128, iv)
        cipher.init(Cipher.DECRYPT_MODE, key(), spec)
        val decrypted = cipher.doFinal(ciphertext)

        return decrypted
    }

    private fun key(): SecretKey {
        return SecretKeySpec(Files.readAllBytes(Path("key.file")), "AES")
    }
}