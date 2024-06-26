package com.torfstack.kayvault.crypto

import org.springframework.stereotype.Service
import sun.awt.Mutex
import java.nio.file.Files
import java.nio.file.OpenOption
import java.nio.file.StandardOpenOption
import java.security.SecureRandom
import java.util.concurrent.locks.ReentrantReadWriteLock
import kotlin.coroutines.CoroutineContext
import kotlin.io.path.Path

@Service
class StaticFileBasedKeyRetriever: KeyRetriever {

    override fun key(): ByteArray {
        return retrieveKey()
    }

    companion object {
        fun retrieveKey(): ByteArray {
            return synchronized(this) {
                if (!Files.exists(Path("key.file"))) {
                    val key = ByteArray(32)
                    SecureRandom().nextBytes(key)
                    Files.write(Path("key.file"), key, StandardOpenOption.CREATE_NEW)
                }
                Files.readAllBytes(Path("key.file"))
            }
        }
    }
}