<?php
namespace functional;

class Simple5Test extends FunctionalTestCase
{
    public function testSomeActivitySimulation(): void
    {
        sleep(7);
        $this->assertTrue(
            true
        );
    }

}