import com.diffplug.gradle.spotless.SpotlessExtension

plugins {
    java
    application
}

java {
    sourceCompatibility = JavaVersion.VERSION_1_8
    targetCompatibility = JavaVersion.VERSION_1_8
}

tasks.register<Exec>("generate") {
    commandLine(
        "buf",
        "generate",
    )
}

tasks.withType<JavaCompile> {
    dependsOn("generate")
}

repositories {
    mavenCentral()
}

dependencies {
    implementation(libs.annotation.api)
    implementation(libs.protobuf.java)
    implementation(libs.protovalidate)

    implementation(platform(libs.grpc.bom))
    implementation(libs.grpc.protobuf)
    implementation(libs.grpc.services)
    implementation(libs.grpc.stub)
    runtimeOnly(libs.grpc.netty.shaded)

    testImplementation(platform(libs.junit.bom))
    testImplementation("org.junit.jupiter:junit-jupiter")
    testRuntimeOnly("org.junit.platform:junit-platform-launcher")
}

buildscript {
    dependencies {
        classpath(libs.spotless)
    }
}

sourceSets {
    main {
        java {
            srcDir(layout.buildDirectory.dir("generated/main/java"))
        }
    }
}

apply(plugin = "com.diffplug.spotless")
configure<SpotlessExtension> {
    java {
        targetExclude("build/generated/main/java/**/*.java")
        palantirJavaFormat()
    }
}

tasks.test {
    useJUnitPlatform()
    testLogging {
        events("passed", "skipped", "failed")
    }
}

application {
    mainClass.set("invoice.v1.InvoiceServer")
}
